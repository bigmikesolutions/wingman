package pkg

import (
	"github.com/bigmikesolutions/wingman/pkg/cqrs"
	"github.com/bigmikesolutions/wingman/pkg/provider"
)

func NewCqrsConfig(providers ...provider.Factory) (cqrs.Config, error) {
	cqrsCfg := cqrs.NewConfig()
	for _, prov := range providers {
		provCfg, err := prov()
		if err != nil {
			return cqrs.Config{}, err
		}

		cqrsCfg, err = cqrs.Merge(cqrsCfg, provCfg)
		if err != nil {
			return cqrs.Config{}, err
		}
	}
	return cqrsCfg, nil
}
