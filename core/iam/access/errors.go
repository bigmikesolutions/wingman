package access

import "github.com/pkg/errors"

type UnauthorizedError struct{ err error }

func NewUnauthorizedError(reason string) error {
	return &UnauthorizedError{errors.Errorf("unauthorized: %s", reason)}
}

func NewUnauthorizedErrorDetails(err error, reason string) error {
	return &UnauthorizedError{errors.Wrapf(err, "unauthorized: %s", reason)}
}

func (e *UnauthorizedError) Error() string {
	return e.err.Error()
}

func (e *UnauthorizedError) Unwrap() error {
	return e.err
}

func (e *UnauthorizedError) Is(target error) bool {
	_, ok := target.(*UnauthorizedError)
	return ok
}
