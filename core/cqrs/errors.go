package cqrs

type HandlerTypeMismatchError struct{ err error }

type MissingHandlerError struct{ err error }

func NewMissingHandlerError(err error) error {
	return &MissingHandlerError{err}
}

func (e *MissingHandlerError) Error() string {
	return e.err.Error()
}

func (e *MissingHandlerError) Unwrap() error {
	return e.err
}

func (e *MissingHandlerError) Is(target error) bool {
	_, ok := target.(*MissingHandlerError)
	return ok
}

func NewHandlerTypeMismatchError(err error) error {
	return &HandlerTypeMismatchError{err}
}

func (e *HandlerTypeMismatchError) Error() string {
	return e.err.Error()
}

func (e *HandlerTypeMismatchError) Unwrap() error {
	return e.err
}

func (e *HandlerTypeMismatchError) Is(target error) bool {
	_, ok := target.(*MissingHandlerError)
	return ok
}
