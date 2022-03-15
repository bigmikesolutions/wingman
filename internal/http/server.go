package http

import (
	"net/http"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/bigmikesolutions/wingman/plugins/k8s"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ctrl struct {
	k8s *k8s.Cluster
}

func NewRouter() (http.Handler, error) {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	k8s, err := k8s.NewCluster()
	if err != nil {
		return nil, err
	}
	ctrl := &ctrl{k8s}
	r.Get("/k8s/pods", ctrl.kubernetesPodsGet)
	return r, nil
}

func (c *ctrl) kubernetesPodsGet(w http.ResponseWriter, r *http.Request) {
	pods, err := c.k8s.GetPods(r.Context(), "cpo", v1.ListOptions{})
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	for _, p := range pods.Items {
		w.Write([]byte(p.Name))
		w.Write([]byte(", "))
	}
}
