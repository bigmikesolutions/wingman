package provider

import (
	"fmt"
	"strings"

	"github.com/bigmikesolutions/wingman/pkg/cqrs"
)

const (
	ProviderGetResourceQueryType cqrs.QueryType = "ProviderGetResource"
)

type ProviderGetResourceQuery struct {
	ProviderID ProviderID
	Path       []string
	Params     map[string][]string
	Query      map[string][]string
}

func (c ProviderGetResourceQuery) GetType() cqrs.QueryType {
	return GetProviderGetResourceQueryType(c.ProviderID, c.Path...)
}

func GetProviderGetResourceQueryType(providerID ProviderID, path ...string) cqrs.QueryType {
	return fmt.Sprintf("%s.%s.%s", providerID, strings.Join(path, "."), ProviderGetResourceQueryType)
}
