package k8s

import (
	"github.com/bigmikesolutions/wingman/pkg/provider"
)

const (
	ProviderName     provider.ProviderID = "k8s"
	PodsResourceType                     = "pods"
)

var AliasNamespace = []string{"ns", "namespace"}
