package env

import (
	"net/http"

	"github.com/bigmikesolutions/wingman/server/auth"
)

const (
	A10NHeaderEnvToken = "X-A10N-Env-Token"
)

type (
	a10nService interface {
		Validate(token string) (*auth.Token, error)
	}
)

func SessionReader(a10n a10nService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			token := r.Header.Get(A10NHeaderEnvToken)
			if token != "" {
				token, err := a10n.Validate(token)
				if err != nil {
					w.WriteHeader(http.StatusForbidden)
					_, _ = w.Write([]byte("invalid env session token"))
					return

				}

				ctx = CtxWithSession(ctx, Session{
					ValidTill: token.ExpiresAt,
				})

				if err := ValidateSession(ctx); err != nil {
					w.WriteHeader(http.StatusForbidden)
					_, _ = w.Write([]byte(err.Error()))
					return
				}
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
