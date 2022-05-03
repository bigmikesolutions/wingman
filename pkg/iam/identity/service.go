package identity

import (
	"context"
)

type AuthService interface {
	IsValid(ctx context.Context, s UserSession) (bool, error)
}

type UserService interface {
	SignIn(ctx context.Context, login string, pass []byte) (UserSession, error)
}

type InMemoryAuthService struct {
	Users map[string]User
}

func (i InMemoryAuthService) IsValid(ctx context.Context, s UserSession) (bool, error) {
	return true, nil
}
