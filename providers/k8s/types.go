package k8s

import "github.com/bigmikesolutions/wingman/pkg/iam"

const (
	ProviderName iam.ProviderID = "k8s"
)

var AliasNamespace = []string{"ns", "namespace"}
