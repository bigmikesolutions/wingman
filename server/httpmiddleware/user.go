package httpmiddleware

import (
	"net/http"
	"strings"

	"github.com/bigmikesolutions/wingman/server/a10n"
)

const (
	headerUserID       = "X-Forwarded-User"
	headerUserName     = "X-Forwarded-Preferred-Username"
	headerUserEmail    = "X-Forwarded-Email"
	headerUserRoles    = "X-A10-Roles"
	userRolesSeparator = ","
)

// UserIdentity reads user info from X headers and puts it into request context.
func UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := a10n.UserIdentity{
			UserID:   r.Header.Get(headerUserID),
			Email:    r.Header.Get(headerUserEmail),
			UserName: r.Header.Get(headerUserName),
		}

		roles := r.Header.Get(headerUserRoles)
		if roles != "" {
			user.Roles = strings.Split(roles, userRolesSeparator)
		}

		ctx := a10n.WithUserIdentity(r.Context(), user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
