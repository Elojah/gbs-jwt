package main

import (
	"context"

	gerrors "github.com/elojah/gbs-jwt/pkg/errors"
	"github.com/elojah/gbs-jwt/pkg/jwt"
	"github.com/gogo/protobuf/types"
	"github.com/gogo/status"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
)

func (h *handler) List(ctx context.Context, req *types.Empty) (*jwt.SecretList, error) {
	logger := log.With().Str("method", "list").Logger()

	if req == nil {
		return &jwt.SecretList{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	result, err := h.jwt.ListSecret(ctx)
	if err != nil {
		// TODO switch status code error depending on error
		return &jwt.SecretList{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &result, nil
}
