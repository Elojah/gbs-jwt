package jwt

import (
	"context"

	gerrors "github.com/elojah/gbs-jwt/pkg/errors"
	gojwt "github.com/golang-jwt/jwt"
)

type App interface {
	CreateSecret(context.Context, Secret) (Secret, error)
	UpdateSecret(context.Context, Secret) (Secret, error)
	DeleteSecret(context.Context, Secret) (Secret, error)
	ListSecret(context.Context) (SecretList, error)
}

// Basic CRUD store
type Store interface {
	Create(context.Context, string, string) error
	Get(context.Context, string) (string, error)
	Update(context.Context, string, string) error
	Delete(context.Context, string) (string, error)
	ListAll(context.Context) (map[string]string, error)
}

func (s Secret) JWT(key string) (string, error) {
	claims := gojwt.MapClaims{}

	for k, v := range s.Claims {
		claims[k] = v
	}

	ss, err := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims).SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func NewSecret(name string, v string, key string) (Secret, error) {
	token, err := gojwt.Parse(v, func(t *gojwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	var validationErr error

	if err != nil {
		ve, ok := err.(*gojwt.ValidationError)
		if !ok || ve.Errors&(gojwt.ValidationErrorExpired|gojwt.ValidationErrorNotValidYet) == 0 {
			return Secret{}, err
		}

		// ! Non fatal error
		validationErr = gerrors.ErrJWTValidation{Name: name, Original: err}
	}

	claims, ok := token.Claims.(gojwt.MapClaims)
	if !ok {
		return Secret{}, gerrors.ErrInvalidFormat{Index: name}
	}

	secret := Secret{
		Name:   name,
		Claims: make(map[string]string, len(claims)),
	}

	for k, v := range claims {
		s, ok := v.(string)
		if !ok {
			continue
		}

		secret.Claims[k] = s
	}

	return secret, validationErr
}
