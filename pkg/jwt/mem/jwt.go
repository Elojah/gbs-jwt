package mem

import (
	"context"

	gerrors "github.com/elojah/gbs-jwt/pkg/errors"
)

func (s *Store) Create(ctx context.Context, key string, value string) error {
	s.Secrets[key] = value

	return nil
}

func (s *Store) Get(ctx context.Context, key string) (string, error) {
	v, ok := s.Secrets[key]
	if !ok {
		return "", gerrors.ErrNotFound{Index: key}
	}

	return v, nil
}

func (s *Store) Update(ctx context.Context, key string, value string) error {
	if _, ok := s.Secrets[key]; !ok {
		return gerrors.ErrNotFound{Index: key}
	}

	s.Secrets[key] = value

	return nil
}

func (s *Store) Delete(ctx context.Context, key string) (string, error) {
	v, ok := s.Secrets[key]
	if !ok {
		return "", gerrors.ErrNotFound{Index: key}
	}

	delete(s.Secrets, key)

	return v, nil
}

func (s *Store) ListAll(ctx context.Context) (map[string]string, error) {
	return s.Secrets, nil
}
