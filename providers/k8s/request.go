package k8s

import (
	"strings"

	"github.com/bigmikesolutions/wingman/pkg/iam"
)

func GetRequestNamespace(request *iam.GetResourceRequest) string {
	for _, alias := range AliasNamespace {
		ns, found := request.Query[alias]
		if found {
			return strings.Join(ns, ",")
		}
	}
	return ""
}

func GetNamespace(query map[string][]string) string {
	for _, alias := range AliasNamespace {
		ns, found := query[alias]
		if found {
			return strings.Join(ns, ",")
		}
	}
	return ""
}
