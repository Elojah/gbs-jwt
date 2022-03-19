package app

import (
	"context"
	"errors"
	"strconv"
	"time"

	gerrors "github.com/elojah/gbs-jwt/pkg/errors"
	"github.com/elojah/gbs-jwt/pkg/jwt"
	"github.com/hashicorp/go-multierror"
)

var _ jwt.App = (*App)(nil)

type App struct {
	Store jwt.Store

	Key        string
	DefaultExp time.Duration
}

func (a *App) CreateSecret(ctx context.Context, secret jwt.Secret) (jwt.Secret, error) {
	if _, err := a.Store.Get(ctx, secret.Name); err != nil {
		if !errors.As(err, &gerrors.ErrNotFound{}) {
			return jwt.Secret{}, err
		}
		// secret not found, valid path
	} else {
		return jwt.Secret{}, gerrors.ErrAlreadyExist{Index: secret.Name}
	}

	if secret.Claims == nil {
		secret.Claims = make(map[string]string)
	}

	// if no expiration is set, set default here
	if _, ok := secret.Claims["exp"]; !ok {
		secret.Claims["exp"] = strconv.Itoa(int(time.Now().Add(a.DefaultExp).Unix()))
	}

	s, err := secret.JWT(a.Key)
	if err != nil {
		return jwt.Secret{}, err
	}

	if err := a.Store.Create(ctx, secret.Name, s); err != nil {
		return jwt.Secret{}, err
	}

	return secret, nil
}

func (a *App) UpdateSecret(ctx context.Context, secret jwt.Secret) (jwt.Secret, error) {
	if _, err := a.Store.Get(ctx, secret.Name); err != nil {
		return jwt.Secret{}, err
	}

	// if no expiration is set, set default here
	if _, ok := secret.Claims["exp"]; !ok {
		secret.Claims["exp"] = strconv.Itoa(int(time.Now().Add(a.DefaultExp).Unix()))
	}

	s, err := secret.JWT(a.Key)
	if err != nil {
		return jwt.Secret{}, err
	}

	if err := a.Store.Update(ctx, secret.Name, s); err != nil {
		return jwt.Secret{}, err
	}

	return secret, nil
}

func (a *App) DeleteSecret(ctx context.Context, secret jwt.Secret) (jwt.Secret, error) {
	token, err := a.Store.Delete(ctx, secret.Name)
	if err != nil {
		return jwt.Secret{}, err
	}

	result, err := jwt.NewSecret(secret.Name, token, a.Key)
	if err != nil {
		if errors.As(err, &gerrors.ErrJWTValidation{}) {
			// ok, still display it
		} else {
			return jwt.Secret{}, err
		}
	}

	return result, nil
}

func (a *App) ListSecret(ctx context.Context) (jwt.SecretList, error) {
	ss, err := a.Store.ListAll(ctx)
	if err != nil {
		return jwt.SecretList{}, err
	}

	var merr *multierror.Error

	var result jwt.SecretList

	// preallocate secrets
	result.Secrets = make([]jwt.Secret, 0, len(ss))

	for k, v := range ss {
		secret, err := jwt.NewSecret(k, v, a.Key)
		if err != nil {
			if errors.As(err, &gerrors.ErrJWTValidation{}) {
				// ok, still list it
			} else {
				merr = multierror.Append(merr, err)

				continue
			}
		}

		result.Secrets = append(result.Secrets, secret)
	}

	if err := merr.ErrorOrNil(); err != nil {
		return jwt.SecretList{}, err
	}

	return result, nil
}

func (a *App) RefreshSecret(ctx context.Context) (jwt.SecretList, error) {
	ss, err := a.Store.ListAll(ctx)
	if err != nil {
		return jwt.SecretList{}, err
	}

	var merr *multierror.Error

	var result jwt.SecretList

	// preallocate secrets
	result.Secrets = make([]jwt.Secret, 0, len(ss))

	for k, v := range ss {
		secret, err := jwt.NewSecret(k, v, a.Key)
		if err != nil {
			if errors.As(err, &gerrors.ErrJWTValidation{}) {

				// refresh token
				secret.Claims["exp"] = strconv.Itoa(int(time.Now().Add(a.DefaultExp).Unix()))
				token, err := secret.JWT(a.Key)
				if err != nil {
					merr = multierror.Append(merr, err)

					continue
				}

				// update refreshed token
				if err := a.Store.Update(ctx, secret.Name, token); err != nil {
					merr = multierror.Append(merr, err)

					continue
				}
			} else {
				merr = multierror.Append(merr, err)

				continue
			}

			result.Secrets = append(result.Secrets, secret)
		}
	}

	if err := merr.ErrorOrNil(); err != nil {
		return jwt.SecretList{}, err
	}

	return result, nil
}
