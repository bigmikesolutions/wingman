package httpmiddleware

import "net/http"

const headerXForwardedUri = "X-Forwarded-Uri"

func RedirectProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		originalPath := r.Header.Get(headerXForwardedUri)
		if originalPath != "" {
			r.URL.Path = originalPath
		}
		next.ServeHTTP(w, r)
	})
}
