package errors

import "fmt"

type ErrNotFound struct {
	Index string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("value not found for index %s", e.Index)
}

type ErrAlreadyExist struct {
	Index string
}

func (e ErrAlreadyExist) Error() string {
	return fmt.Sprintf("value already exist for index %s", e.Index)
}

type ErrInvalidFormat struct {
	Index string
}

func (e ErrInvalidFormat) Error() string {
	return fmt.Sprintf("value wrongly formatted for index %s", e.Index)
}
