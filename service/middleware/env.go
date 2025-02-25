package middleware

import (
	"net/http"

	"github.com/bigmikesolutions/wingman/service/auth"
)

const (
	A10NHeaderEnvToken = "X-A10N-Env-Token"
)

type (
	a10nService interface {
		Validate(token string) (*auth.Token, error)
	}
)

func EnvSession(a10n a10nService, next http.Handler) http.Handler {
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

			ctx = auth.CtxWithEnvSession(ctx, auth.EnvSession{
				ValidTill: token.ExpiresAt,
			})

			if err := auth.ValidateEnvSession(ctx); err != nil {
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write([]byte(err.Error()))
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
