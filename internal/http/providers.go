package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/bigmikesolutions/wingman/pkg/cqrs"

	"github.com/bigmikesolutions/wingman/pkg/provider"
)

type ProviderCtrl struct {
	path string
	cqrs *cqrs.CQRS
}

func NewController(path string, cqrs *cqrs.CQRS) *ProviderCtrl {
	return &ProviderCtrl{
		path: path,
		cqrs: cqrs,
	}
}

func (c *ProviderCtrl) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	reqPath := strings.Split(
		strings.TrimRight(
			strings.TrimLeft(request.URL.Path, c.path),
			"/",
		),
		"/")
	switch request.Method {
	case http.MethodGet:
		if reqPath[0] == "" {
			c.list(writer, request)
		} else {
			c.getProviderResource(writer, request, provider.ID(reqPath[0]), reqPath[1:]...)
		}

	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c *ProviderCtrl) getProviderResource(w http.ResponseWriter, r *http.Request, providerID provider.ID, resourcePath ...string) {
	result, err := c.cqrs.ExecuteQuery(r.Context(), &provider.QueryGetResource{
		ProviderID: providerID,
		Path:       resourcePath,
		Params:     r.Header,
		Query:      r.URL.Query(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("couldn't marshal resource: %s", err.Error())))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

func (c *ProviderCtrl) list(writer http.ResponseWriter, request *http.Request) {
	// TODO: implement
	providers := "not supported"
	json, err := json.Marshal(providers)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf("couldn't marshal providers list: %s", err.Error())))
	} else {
		writer.WriteHeader(http.StatusOK)
		writer.Write(json)
	}
}
