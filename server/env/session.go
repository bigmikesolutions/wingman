package env

import (
	"context"
	"errors"
	"time"
)

type (
	ctxKey struct{}

	Session struct {
		ValidTill time.Time
	}
)

var (
	ctxKeyValue ctxKey

	ErrNoActiveEnvSession = errors.New("no active environment session")
)

func ValidateSession(ctx context.Context) error {
	s, ok := GetSession(ctx)

	if !ok {
		return ErrNoActiveEnvSession
	}

	if !s.ValidTill.After(time.Now()) {
		return errors.New("env session has expired")
	}
	return nil
}

func WithSession(ctx context.Context, session Session) context.Context {
	return context.WithValue(ctx, ctxKeyValue, session)
}

func GetSession(ctx context.Context) (Session, bool) {
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
