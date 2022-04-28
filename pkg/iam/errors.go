package iam

import "github.com/pkg/errors"

type UnauthenticatedError struct{ err error }

func newUnauthenticatedError(reason string) error {
	return &UnauthenticatedError{errors.Errorf("unauthenticated user: %s", reason)}
}

func newUnauthenticatedErrorDetails(err error, reason string) error {
	return &UnauthenticatedError{errors.Wrapf(err, "unauthenticated user: %s", reason)}
}

func (e *UnauthenticatedError) Error() string {
	return e.err.Error()
}

func (e *UnauthenticatedError) Unwrap() error {
	return e.err
}

func (e *UnauthenticatedError) Is(target error) bool {
	_, ok := target.(*UnauthenticatedError)
	return ok
}
