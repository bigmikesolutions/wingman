package iam

import (
	"net/http"
)

type HttpAuthCtrl struct {
	iam identity.AuthService
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
