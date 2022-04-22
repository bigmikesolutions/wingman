package iam

import (
	"net/http"

	"github.com/bigmikesolutions/wingman/pkg/iam"
)

type HttpAuthCtrl struct {
	iam iam.AuthService
}

func (c *HttpAuthCtrl) AuthUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (c *HttpAuthCtrl) AuthRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
