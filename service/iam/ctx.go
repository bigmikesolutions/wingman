package iam

import (
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

var userCtxKey = &contextKey{}

func CtxWithUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

func CtxUser(ctx context.Context) (User, bool) {
	if u, ok := ctx.Value(userCtxKey).(User); ok {
		return u, true
	} else {
		return User{}, false
	}
}
