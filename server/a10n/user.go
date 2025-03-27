package a10n

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

type (
	contextKey struct{}

	Role           = string
	OrganisationID = string
	UserID         = string

	UserIdentity struct {
		UserID   UserID
		Email    string
		UserName string
		OrgID    OrganisationID
		Roles    []Role
	}
)

const (
	AdminRead      Role = "admin-read"
	AdminWrite     Role = "admin-write"
	ManagerRead    Role = "manager-read"
	ManagerWrite   Role = "manager-write"
	DeveloperRead  Role = "dev-read"
	DeveloperWrite Role = "dev-write"
)

var (
	ctxKey                  = &contextKey{}
	ErrUserNotAuthenticated = errors.New("user not authenticated")
	ErrUserNotAuthorized    = errors.New("user not authorized")
)

func deepClone(u UserIdentity) UserIdentity {
	c := make([]Role, len(u.Roles))
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

func UserAuthenticated(ctx context.Context) error {
	_, err := GetIdentity(ctx)
	return err
}

func UserAuthorized(ctx context.Context, roles ...Role) error {
	u, err := GetIdentity(ctx)
	if err != nil {
		return err
	}

	return u.ContainsRole(roles...)
}

func (u *UserIdentity) ContainsRole(roles ...Role) error {
	for _, r := range roles {
		for _, role := range u.Roles {
			if role == r {
				return nil
			}
		}
	}

	return fmt.Errorf("%w - missing one scopes: %s",
		ErrUserNotAuthorized,
		strings.Join(roles, ","),
	)
}
