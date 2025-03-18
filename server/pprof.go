package server

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi/v5"
)

func pprofRouter() http.Handler {
	r := chi.NewRouter()
	r.HandleFunc("/", pprof.Index)

	r.HandleFunc("/cmdline", pprof.Cmdline)
	r.HandleFunc("/profile", pprof.Profile)
	r.HandleFunc("/symbol", pprof.Symbol)
	r.HandleFunc("/trace", pprof.Trace)

	r.Handle("/cmdline", pprof.Handler("block"))
	r.Handle("/goroutine", pprof.Handler("goroutine"))
	r.Handle("/heap", pprof.Handler("heap"))
	r.Handle("/threadcreate", pprof.Handler("threadcreate"))

	return r
}
