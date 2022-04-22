package providers

import (
	"sync"

	"github.com/bigmikesolutions/wingman/pkg/iam"
	"github.com/bigmikesolutions/wingman/providers/k8s"
	"github.com/pkg/errors"
)

type ProviderFactory = func() (iam.ResourceProvider, error)

type Providers struct {
	mux            sync.RWMutex
	providersList  []iam.ProviderID
	providers      map[iam.ProviderID]ProviderFactory
	providersCache map[iam.ProviderID]iam.ResourceProvider
}

func NewProviders() *Providers {
	providers := map[iam.ProviderID]ProviderFactory{
		k8s.ProviderName: k8s.NewProvider,
	}
	providersList := make([]iam.ProviderID, 0)
	for id := range providers {
		providersList = append(providersList, id)
	}
	return &Providers{
		providersCache: make(map[iam.ProviderID]iam.ResourceProvider),
		providers:      providers,
		providersList:  providersList,
	}
}

func (p Providers) Find(id iam.ProviderID) (iam.ResourceProvider, error) {
	// check cache
	p.mux.RLock()
	provider, _ := p.providersCache[id]
	p.mux.RUnlock()
	if provider != nil {
		return provider, nil
	}
	// check providers
	factory, factoryFound := p.providers[id]
	if !factoryFound {
		return nil, errors.Errorf("provider %s not found", id)
	}
	// create new provider
	provider, err := factory()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create provider %s", id)
	}
	p.mux.Lock()
	p.providersCache[id] = provider
	p.mux.Unlock()
	return provider, nil
}

func (p Providers) List() []iam.ProviderID {
	return p.providersList
}
