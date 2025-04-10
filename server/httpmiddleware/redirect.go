package httpmiddleware

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

const (
	headerXForwardedUri    = "X-Forwarded-Uri"
	headerXForwardedMethod = "X-Forwarded-Method"
)

func RedirectProxy(next http.Handler) http.Handler {
	logger := log.Logger
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		originalPath := r.Header.Get(headerXForwardedUri)
		if originalPath != "" {
			r.URL.Path = originalPath
		}

		originalMethod := r.Header.Get(headerXForwardedMethod)
		if originalMethod != "" && originalPath != r.Method {
			logger.Warn().
				Str("method", r.Method).
				Str("original_method", originalMethod).
				Str("original_path", originalPath).
				Msg("redirect proxy request method does not match - request body might be lost")
			r.Method = originalMethod
		}

		next.ServeHTTP(w, r)
	})
}
