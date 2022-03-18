package mem

import (
	"context"

	"github.com/elojah/gbs-jwt/pkg/jwt"
)

var _ jwt.Store = (*Store)(nil)

// Store in memory for test/demonstration purpose only.
type Store struct {
	Secrets map[string]string
}

func NewStore(ctx context.Context) (Store, error) {
	return Store{
		Secrets: make(map[string]string),
	}, nil
}
