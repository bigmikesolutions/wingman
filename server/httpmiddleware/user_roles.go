package httpmiddleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const (
	headerXForwardedAccessToken = "X-Forwarded-Access-Token"
	keyRoles                    = "roles"
)

// UserRoles extract roles from JWT token and saves it as one of X headers.
func UserRoles(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get(headerXForwardedAccessToken)

		if tokenHeader != "" {
			tokenClaims := jwt.MapClaims{}
			_, _, err := jwt.NewParser().ParseUnverified(tokenHeader, &tokenClaims)
			if err != nil {
				http.Error(w, "jwt decode: "+err.Error(), http.StatusUnauthorized)
				return
			}

			roles := extractRoles(tokenClaims)
			if len(roles) > 0 {
				r.Header.Set(headerUserRoles, strings.Join(roles, userRolesSeparator))
			}

		}

		next.ServeHTTP(w, r)
	})
}

func extractRoles(m map[string]any) []string {
	roles := make([]string, 0)
	for k, v := range m {
		if k == keyRoles {
			if values, ok := v.([]any); ok {
				for _, r := range values {
					roles = append(roles, fmt.Sprintf("%s", r))
				}
			}
		}

		if m, ok := v.(map[string]any); ok {
			roles = append(roles, extractRoles(m)...)
		}
	}
	return roles
}
