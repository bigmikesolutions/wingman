package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/bigmikesolutions/wingman/pkg/iam"
)

type ProviderCtrl struct {
	path      string
	providers iam.ProviderPlugins
}

func (c *ProviderCtrl) Mount(handler http.Handler) {
}

func (c *ProviderCtrl) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	reqPath := strings.Split(strings.TrimLeft(request.URL.Path, c.path), "/")
	switch request.Method {
	case http.MethodGet:
		if reqPath[0] == "" {
			c.list(writer, request)
		} else {
			c.getProviderResource(writer, request, iam.ProviderID(reqPath[0]), reqPath[1:]...)
		}

	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c *ProviderCtrl) getProviderResource(w http.ResponseWriter, r *http.Request, providerID iam.ProviderID, resourcePath ...string) {
	provider, err := c.providers.Find(providerID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else if provider == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("provider not found: %s", string(providerID))))
	} else {
		resource, err := provider.Provide(&iam.GetResourceRequest{
			Ctx:    r.Context(),
			Path:   resourcePath,
			Params: r.Header,
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("provider resource get error: %s", err.Error())))
		}
		json, err := json.Marshal(resource.GetInfo())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("couldn't marshal resource: %s", err.Error())))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(json)
		}
	}
}

func (c *ProviderCtrl) list(writer http.ResponseWriter, request *http.Request) {
	providers := c.providers.List()
	json, err := json.Marshal(providers)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf("couldn't marshal providers list: %s", err.Error())))
	} else {
		writer.WriteHeader(http.StatusOK)
		writer.Write(json)
	}
}
