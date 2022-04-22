package iam

import "context"

type AuthService interface {
	SignIn(login string, pass []byte) (UserSession, error)
	IsValid(s UserSession) (bool, error)
}

type AccessService interface {
	AccessResource(req UserEnvironmentAccessRequest) (UserEnvironmentSession, error)
}

type GetResourceRequest struct {
	Path   []string
	Params map[string][]string
	Ctx    context.Context
	Query  map[string][]string
}

type ResourceProvider interface {
	Provide(req *GetResourceRequest) (Resource, error)
}

type ProviderPlugins interface {
	List() []ProviderID
	Find(id ProviderID) (ResourceProvider, error)
}
