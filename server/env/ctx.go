package env

import (
	"context"
	"errors"
	"time"
)

type (
	ctxKey struct{}
)

var (
	ctxKeyValue ctxKey

	errNoActiveEnvSession = errors.New("no active environment session")
)

func ValidateSession(ctx context.Context) error {
	s, ok := FromContext(ctx)

	if !ok {
		return errNoActiveEnvSession
	}

	if !s.ValidTill.After(time.Now()) {
		return errors.New("env session has expired")
	}
	return nil
}

func CtxWithSession(ctx context.Context, session Session) context.Context {
	return context.WithValue(ctx, ctxKeyValue, session)
}

func FromContext(ctx context.Context) (Session, bool) {
	v := ctx.Value(ctxKeyValue)
	if v == nil {
		return Session{}, false
	}

	s, ok := v.(Session)
	if !ok {
		return Session{}, false
	}

	return s, true
}
