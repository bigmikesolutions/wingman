package auth

import (
	"net/http"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	headerAuthorization = "Authorization"
	tokenBearer         = "Bearer "
)

func UserIdentity(level zerolog.Level) func(http.Handler) http.Handler {
	logger := log.Logger
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			event := logger.WithLevel(level)

			if event.Enabled() {
				auth := r.Header.Get(headerAuthorization)
				if auth != "" && strings.HasPrefix(auth, tokenBearer) {
					token := strings.Split(auth, " ")[1]
					if token != "" {
						event = event.Str("token", token)
						event.Msg("User identity set")
					}
				}

			}

			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	return f
}
