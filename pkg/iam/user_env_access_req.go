package iam

import "github.com/bigmikesolutions/wingman/pkg/provider"

type UserEnvironmentAccessRequest struct {
	userRequestingAccess   UserSession
	pairingWith            User
	reason                 string
	incident               string
	resourceAccessRequests map[provider.Resource][]provider.AccessType
}

func NewUserEnvironmentAccessRequest(userRequestingAccess UserSession, pairingWith User, reason string, incident string, resourceAccessRequests map[provider.Resource][]provider.AccessType) *UserEnvironmentAccessRequest {
	return &UserEnvironmentAccessRequest{userRequestingAccess: userRequestingAccess, pairingWith: pairingWith, reason: reason, incident: incident, resourceAccessRequests: resourceAccessRequests}
}

func (u UserEnvironmentAccessRequest) UserRequestingAccess() UserSession {
	return u.userRequestingAccess
}

func (u UserEnvironmentAccessRequest) PairingWith() User {
	return u.pairingWith
}

func (u UserEnvironmentAccessRequest) Reason() string {
	return u.reason
}

func (u UserEnvironmentAccessRequest) Incident() string {
	return u.incident
}

func (u UserEnvironmentAccessRequest) ResourceAccessRequests() map[provider.Resource][]provider.AccessType {
	return u.resourceAccessRequests
}
