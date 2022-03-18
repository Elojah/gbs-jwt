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

func (h *handler) Delete(ctx context.Context, req *jwt.Secret) (*types.Empty, error) {
	logger := log.With().Str("method", "delete").Logger()

	if req == nil {
		return &types.Empty{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	_, err := h.jwt.DeleteSecret(ctx, *req)
	if err != nil {
		// TODO switch status code error depending on error
		return &types.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &types.Empty{}, nil
}
