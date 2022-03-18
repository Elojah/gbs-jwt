package main

import (
	"context"

	"github.com/elojah/go-grpc"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	GRPC grpc.Config `json:"grpc"`
	// JWT  jwt.Config  `json:"jwt"`
}

// Populate populates config object reading file and env.
func (c *config) Populate(ctx context.Context, filename string) error {
	return cleanenv.ReadConfig(filename, c)
}
