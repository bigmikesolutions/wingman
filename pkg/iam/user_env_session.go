package iam

type (
	UserEnvironmentSessionID string
)

type UserEnvironmentSession struct {
	id                     UserEnvironmentSessionID
	user                   UserSession
	resourceAccessRequests []ResourceAccessSession
}

func NewUserEnvironmentSession(id UserEnvironmentSessionID, user UserSession, resourceAccessRequests ...ResourceAccessSession) *UserEnvironmentSession {
	return &UserEnvironmentSession{id: id, user: user, resourceAccessRequests: resourceAccessRequests}
}

func (u UserEnvironmentSession) Id() UserEnvironmentSessionID {
	return u.id
}

func (u UserEnvironmentSession) User() UserSession {
	return u.user
}

func (u UserEnvironmentSession) ResourceAccessRequests() []ResourceAccessSession {
	return u.resourceAccessRequests
}
