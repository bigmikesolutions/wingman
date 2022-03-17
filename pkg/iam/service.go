package iam

type Service interface {
	SignIn(login string, pass []byte) (UserSession, error)
	IsValid(s UserSession) (bool, error)
	AccessResource(req UserEnvironmentAccessRequest) (UserEnvironmentSession, error)
	GetResource(id ResourceID) Resource
}
