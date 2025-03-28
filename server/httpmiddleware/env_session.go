package httpmiddleware

import (
	"net/http"

	"github.com/bigmikesolutions/wingman/server/env"
	"github.com/bigmikesolutions/wingman/server/token"
)

const (
	HeaderEnvToken = "X-Forwarded-Env-Token"
)

type (
	a10nService interface {
		Validate(token string) (*token.Token, error)
	}
)

func SessionReader(a10n a10nService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			headerToken := r.Header.Get(HeaderEnvToken)
			if headerToken != "" {
				t, err := a10n.Validate(headerToken)
				if err != nil {
					w.WriteHeader(http.StatusForbidden)
					_, _ = w.Write([]byte("invalid env session token"))
					return

				}

				ctx = env.WithSession(ctx, env.Session{
					ValidTill: t.ExpiresAt,
				})

				if err := env.ValidateSession(ctx); err != nil {
					w.WriteHeader(http.StatusForbidden)
					_, _ = w.Write([]byte(err.Error()))
					return
				}
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
