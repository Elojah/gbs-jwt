package app

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/elojah/gbs-jwt/pkg/jwt"
	"github.com/elojah/gbs-jwt/pkg/jwt/mem"
	gojwt "github.com/golang-jwt/jwt"
)

func TestApp(t *testing.T) {
	ctx := context.Background()

	s, err := mem.NewStore(ctx)
	if err != nil {
		t.Fatal(err)

		return
	}

	a := App{
		Store:      &s,
		Key:        "randomkey",
		DefaultExp: time.Hour,
	}

	// Create
	// List
	// Update
	// List
	// Delete
	// List

	type mock struct {
		Create     []jwt.Secret
		FirstList  jwt.SecretList
		Update     []jwt.Secret
		SecondList jwt.SecretList
		Delete     []jwt.Secret
		ThirdList  jwt.SecretList
	}

	for _, d := range []mock{
		{
			Create: []jwt.Secret{{
				Name: "test0",
				Claims: map[string]string{
					"name": "carapace",
					"exp":  "57823653896539",
				},
			}},
			FirstList: jwt.SecretList{
				Secrets: []jwt.Secret{{
					Name: "test0",
					Claims: map[string]string{
						"name": "carapace",
						"exp":  "57823653896539",
					},
				}},
			},
			Update: []jwt.Secret{{
				Name: "test0",
				Claims: map[string]string{
					"name": "carapouce",
					"exp":  "57823653896539",
				},
			}},
			SecondList: jwt.SecretList{
				Secrets: []jwt.Secret{{
					Name: "test0",
					Claims: map[string]string{
						"name": "carapouce",
						"exp":  "57823653896539",
					},
				}},
			},
			Delete: []jwt.Secret{{
				Name: "test0",
			}},
			ThirdList: jwt.SecretList{
				Secrets: []jwt.Secret{},
			},
		},
	} {
		// Create
		for _, s := range d.Create {
			_, err := a.CreateSecret(ctx, s)
			if err != nil {
				t.Fatal(err)

				return
			}
		}

		actual, err := a.ListSecret(ctx)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(actual, d.FirstList) {
			fmt.Println(actual, d.FirstList)
			t.Fatal("first list not equal")

			return
		}

		// Update
		for _, s := range d.Update {
			_, err := a.UpdateSecret(ctx, s)
			if err != nil {
				t.Fatal(err)

				return
			}
		}

		actual, err = a.ListSecret(ctx)
		if err != nil {
			t.Fatal(err)

			return
		}

		if !reflect.DeepEqual(actual, d.SecondList) {
			t.Fatal("first list not equal")

			return
		}

		// Delete
		for _, s := range d.Delete {
			_, err := a.DeleteSecret(ctx, s)
			if err != nil {
				t.Fatal(err)

				return
			}
		}

		actual, err = a.ListSecret(ctx)
		if err != nil {
			t.Fatal(err)

			return
		}

		if !reflect.DeepEqual(actual, d.ThirdList) {
			t.Fatal("first list not equal")

			return
		}
	}
}

func TestJWTValidity(t *testing.T) {
	ctx := context.Background()

	s, err := mem.NewStore(ctx)
	if err != nil {
		t.Fatal(err)

		return
	}

	a := App{
		Store:      &s,
		Key:        "randomkey",
		DefaultExp: time.Hour,
	}

	if _, err := a.CreateSecret(ctx, jwt.Secret{
		Name: "test",
		Claims: map[string]string{
			"key": "value",
		},
	}); err != nil {
		t.Fatal(err)

		return
	}

	ss, err := s.Get(ctx, "test")
	if err != nil {
		t.Fatal(err)

		return
	}

	token, err := gojwt.Parse(ss, func(t *gojwt.Token) (interface{}, error) {
		return []byte(a.Key), nil
	})
	if err != nil {
		t.Fatal(err)

		return
	}

	claims, ok := token.Claims.(gojwt.MapClaims)
	if !ok {
		t.Fatal("failed to read token claims")

		return
	}

	if v, ok := claims["key"]; !ok || v != "value" {
		t.Fatal("claim invalid")

		return
	}
}
