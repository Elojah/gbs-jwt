package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	secretsgrpc "github.com/elojah/gbs-jwt/cmd/secrets/grpc"
	jwtapp "github.com/elojah/gbs-jwt/pkg/jwt/app"
	jwtmem "github.com/elojah/gbs-jwt/pkg/jwt/mem"
	ggrpc "github.com/elojah/go-grpc"
	glog "github.com/elojah/go-log"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog/log"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"
)

const (
	// Time allocated for init phase (connections + setup).
	initTO = 30 * time.Second
)

var version string

type closer interface {
	Close(context.Context) error
}

type closers []closer

func (cs closers) Close(ctx context.Context) error {
	var result *multierror.Error

	for _, c := range cs {
		result = multierror.Append(result, c.Close(ctx))
	}

	return result.ErrorOrNil()
}

// run services.
func run(prog string, filename string) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, initTO)
	// no need for defer cancel
	_ = cancel

	var cs []closer

	logs := glog.Service{}
	if err := logs.Dial(ctx, glog.Config{}); err != nil {
		fmt.Println("failed to dial logger")

		return
	}

	cs = append(cs, &logs)

	log.Logger = logs.With().Str("version", version).Str("exe", prog).Logger()

	// read config
	cfg := config{}
	if err := cfg.Populate(ctx, filename); err != nil {
		log.Error().Err(err).Msg("failed to read config")

		return
	}

	// store
	jwtStore, err := jwtmem.NewStore(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to init mem store")

		return
	}

	// app definition
	jwtApp := jwtapp.App{
		Store: &jwtStore,

		// TODO set from config
		Key:        "somerandomencryptedstr",
		DefaultExp: time.Duration(1 * time.Hour),
	}

	// init handler
	h := handler{
		jwt: &jwtApp,
	}

	// init grpc api server
	grpcsecrets := ggrpc.Service{}
	grpcsecrets.Register = func() {
		reflection.Register(grpcsecrets.Server)
		secretsgrpc.RegisterSecretsServer(grpcsecrets.Server, &h)
	}

	if err := grpcsecrets.Dial(ctx, cfg.GRPC); err != nil {
		log.Error().Err(err).Msg("failed to dial grpc")

		return
	}

	cs = append(cs, &grpcsecrets)

	go func() {
		if err := grpcsecrets.Serve(ctx); err != nil {
			log.Error().Err(err).Msg("failed to serve grpc")
		}
	}()

	log.Info().Msg("secrets up")

	// listen for signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for sig := range c {
		switch sig {
		case syscall.SIGHUP:
			fallthrough
		case syscall.SIGINT:
			fallthrough
		case syscall.SIGTERM:
			if err := closers(cs).Close(ctx); err != nil {
				fmt.Printf("error closing service: %s\n", err.Error())

				return
			}

			cancel()

			fmt.Println("successfully closed service")

			return
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 2 { // nolint: gomnd
		fmt.Printf("Usage: ./%s configfile\n", args[0])

		return
	}

	run(args[0], args[1])
}
