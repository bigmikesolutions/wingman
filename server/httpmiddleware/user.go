package httpmiddleware

import (
	"net/http"
	"strings"

	"github.com/bigmikesolutions/wingman/server/a10n"
)

const (
	HeaderUserID    = "X-Forwarded-User"
	HeaderUserName  = "X-Forwarded-Preferred-Username"
	HeaderUserEmail = "X-Forwarded-Email"
	HeaderOrgID     = "X-Forwarded-Organisation"
	HeaderUserRoles = "X-Forwarded-Roles"
	RolesSeparator  = ","
)

// UserIdentity reads user info from X headers and puts it into request context.
func UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := a10n.UserIdentity{
			UserID:   r.Header.Get(HeaderUserID),
			Email:    r.Header.Get(HeaderUserEmail),
			UserName: r.Header.Get(HeaderUserName),
			OrgID:    r.Header.Get(HeaderOrgID),
		}

		roles := r.Header.Get(HeaderUserRoles)
		if roles != "" {
			user.Roles = strings.Split(roles, RolesSeparator)
		}

		ctx := a10n.WithUserIdentity(r.Context(), user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
