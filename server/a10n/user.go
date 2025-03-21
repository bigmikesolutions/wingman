package a10n

import (
	"context"
	"errors"
)

type (
	contextKey struct{}

	UserIdentity struct {
		UserID   string
		Email    string
		UserName string
		Roles    []string
	}
)

var (
	ctxKey                  = &contextKey{}
	ErrUserNotAuthenticated = errors.New("user not authenticated")
)

func deepClone(u UserIdentity) UserIdentity {
	c := make([]string, len(u.Roles))
	copy(c, u.Roles)
	u.Roles = c
	return u
}

func WithUserIdentity(ctx context.Context, user UserIdentity) context.Context {
	return context.WithValue(ctx, ctxKey, user)
}

func GetIdentity(ctx context.Context) (UserIdentity, error) {
	if u, ok := ctx.Value(ctxKey).(UserIdentity); ok {
		return deepClone(u), nil
	} else {
		return UserIdentity{}, ErrUserNotAuthenticated
	}
}
