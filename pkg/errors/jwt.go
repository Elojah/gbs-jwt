package errors

import "fmt"

type ErrJWTValidation struct {
	Name     string
	Original error
}

func (e ErrJWTValidation) Error() string {
	return fmt.Sprintf("validation error for token %s", e.Name)
}
