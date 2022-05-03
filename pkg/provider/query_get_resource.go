package provider

import (
	"fmt"
	"strings"

	"github.com/bigmikesolutions/wingman/pkg/cqrs"
)

const (
	QueryTypeGetResource cqrs.QueryType = "ProviderGetResource"
)

type QueryGetResource struct {
	ProviderID ProviderID
	Path       []string
	Params     map[string][]string
	Query      map[string][]string
}

func (c QueryGetResource) GetType() cqrs.QueryType {
	return GetProviderGetResourceQueryType(c.ProviderID, c.Path...)
}

func GetProviderGetResourceQueryType(providerID ProviderID, path ...string) cqrs.QueryType {
	return fmt.Sprintf("%s.%s.%s", providerID, strings.Join(path, "."), QueryTypeGetResource)
}
