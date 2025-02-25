package auth

import (
	"context"
	"errors"
	"time"
)

type (
	envCtxKey  struct{}
	EnvSession struct {
		ValidTill time.Time
	}
)

var envSessionKey envCtxKey

func ValidateEnvSession(ctx context.Context) error {
	s, ok := FromContext(ctx)

	if !ok {
		return errors.New("no env session")
	}

	if !s.ValidTill.After(time.Now()) {
		return errors.New("env session has expired")
	}
	return nil
}

func CtxWithEnvSession(ctx context.Context, session EnvSession) context.Context {
	return context.WithValue(ctx, envSessionKey, session)
}

func FromContext(ctx context.Context) (EnvSession, bool) {
	v := ctx.Value(envSessionKey)
	if v == nil {
		return EnvSession{}, false
	}

	s, ok := v.(EnvSession)
	if !ok {
		return EnvSession{}, false
	}

	return s, true
}
