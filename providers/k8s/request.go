package k8s

import (
	"strings"
)

func GetNamespace(query map[string][]string) string {
	for _, alias := range AliasNamespace {
		ns, found := query[alias]
		if found {
			return strings.Join(ns, ",")
		}
	}
	return ""
}
