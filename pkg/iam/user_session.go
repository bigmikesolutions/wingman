package iam

import "time"

type (
	UserSessionID string
)

type UserSession struct {
	id             UserSessionID
	user           User
	expirationDate time.Time
}

func NewUserSession(id UserSessionID, user User, expirationDate time.Time) *UserSession {
	return &UserSession{id: id, user: user, expirationDate: expirationDate}
}

func (u UserSession) Id() UserSessionID {
	return u.id
}

func (u UserSession) User() User {
	return u.user
}

func (u UserSession) ExpirationDate() time.Time {
	return u.expirationDate
}
