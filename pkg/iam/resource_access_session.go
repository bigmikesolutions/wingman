package iam

import "time"

type (
	ResourceAccessSessionID string
)

type ResourceAccessSession struct {
	id             ResourceAccessSessionID
	resource       Resource
	userSession    UserSession
	accessTypes    []AccessType
	expirationDate time.Time
}

func NewResourceAccessSession(id ResourceAccessSessionID, resource Resource, userSession UserSession, expirationDate time.Time, accessTypes ...AccessType) *ResourceAccessSession {
	return &ResourceAccessSession{id: id, resource: resource, userSession: userSession, accessTypes: accessTypes, expirationDate: expirationDate}
}

func (r ResourceAccessSession) Id() ResourceAccessSessionID {
	return r.id
}

func (r ResourceAccessSession) Resource() Resource {
	return r.resource
}

func (r ResourceAccessSession) UserSession() UserSession {
	return r.userSession
}

func (r ResourceAccessSession) AccessTypes() []AccessType {
	return r.accessTypes
}

func (r ResourceAccessSession) ExpirationDate() time.Time {
	return r.expirationDate
}
