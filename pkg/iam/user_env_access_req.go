package iam

type UserEnvironmentAccessRequest struct {
	userRequestingAccess   UserSession
	pairingWith            User
	reason                 string
	incident               string
	resourceAccessRequests map[Resource][]AccessType
}

func NewUserEnvironmentAccessRequest(userRequestingAccess UserSession, pairingWith User, reason string, incident string, resourceAccessRequests map[Resource][]AccessType) *UserEnvironmentAccessRequest {
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

func (u UserEnvironmentAccessRequest) ResourceAccessRequests() map[Resource][]AccessType {
	return u.resourceAccessRequests
}
