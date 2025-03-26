package a10n

import (
	"context"
	"errors"
	"fmt"
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
	ErrUserNotAuthorized    = errors.New("user not authorized")
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

func Authorized(ctx context.Context, roles ...string) error {
	u, err := GetIdentity(ctx)
	if err != nil {
		return err
	}

	return u.HasRoles(roles...)
}

func (u *UserIdentity) HasRoles(roles ...string) error {
	for _, r := range roles {
		found := false

		for _, role := range u.Roles {
			if role == r {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("%w - missing scope: %s", ErrUserNotAuthorized, r)
		}
	}

	return nil
}
