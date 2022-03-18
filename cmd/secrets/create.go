package main

import (
	"context"

	gerrors "github.com/elojah/gbs-jwt/pkg/errors"
	"github.com/elojah/gbs-jwt/pkg/jwt"
	"github.com/gogo/status"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
)

func (h *handler) Create(ctx context.Context, req *jwt.Secret) (*jwt.Secret, error) {
	logger := log.With().Str("method", "create").Logger()

	if req == nil {
		return &jwt.Secret{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	result, err := h.jwt.CreateSecret(ctx, *req)
	if err != nil {
		// TODO switch status code error depending on error
		return &jwt.Secret{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &result, nil
}
