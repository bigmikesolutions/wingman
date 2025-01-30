package iam

import (
	"errors"

	"golang.org/x/net/context"
)

type (
	contextKey struct{}

	UserCtx struct {
		OrgID  OrganisationID
		UserID UserID
		Roles  []UserRoleBinding
	}
)

var (
	userCtxKey              = &contextKey{}
	ErrUserNotAuthenticated = errors.New("user not authenticated")
	ErrUserNotAuthorized    = errors.New("user not authorized")
)

func CtxWithUser(ctx context.Context, user UserCtx) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

func CtxUser(ctx context.Context) (UserCtx, error) {
	if u, ok := ctx.Value(userCtxKey).(UserCtx); ok {
		return u, nil
	} else {
		return UserCtx{}, ErrUserNotAuthenticated
	}
}
