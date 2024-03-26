package mock

import (
	"context"
	"net/http"
	"time"

	"github.com/bigmikesolutions/wingman/pkg/iam/identity"
)

var MockUserSession = identity.NewUserSession(
	"mock-user-id",
	*identity.NewUser("mock-user-id", "mock-user-name"),
	time.Now().Add(24*time.Hour),
)

type InMemoryAuthService struct {
	Users map[string]identity.User
}

func (i InMemoryAuthService) IsValid(ctx context.Context, s identity.UserSession) (bool, error) {
	return true, nil
}

func MockUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := identity.SetUserSession(r.Context(), *MockUserSession)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
