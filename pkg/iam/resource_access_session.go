package iam

import (
	"time"

	"github.com/bigmikesolutions/wingman/pkg/provider"
)

type (
	ResourceAccessSessionID string
)

type ResourceAccessSession struct {
	id             ResourceAccessSessionID
	resource       provider.Resource
	userSession    UserSession
	accessTypes    []provider.AccessType
	expirationDate time.Time
}

func NewResourceAccessSession(id ResourceAccessSessionID, resource provider.Resource, userSession UserSession, expirationDate time.Time, accessTypes ...provider.AccessType) *ResourceAccessSession {
	return &ResourceAccessSession{id: id, resource: resource, userSession: userSession, accessTypes: accessTypes, expirationDate: expirationDate}
}

func (r ResourceAccessSession) Id() ResourceAccessSessionID {
	return r.id
}

func (r ResourceAccessSession) Resource() provider.Resource {
	return r.resource
}

func (r ResourceAccessSession) UserSession() UserSession {
	return r.userSession
}

func (r ResourceAccessSession) AccessTypes() []provider.AccessType {
	return r.accessTypes
}

func (r ResourceAccessSession) ExpirationDate() time.Time {
	return r.expirationDate
}
