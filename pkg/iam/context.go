package iam

import "context"

type ctx string

const (
	contextUserSession ctx = "iam_user_session"
)

func SetUserSession(ctx context.Context, userSession UserSession) context.Context {
	return context.WithValue(ctx, contextUserSession, userSession)
}

func GetUserSession(ctx context.Context) *UserSession {
	if ctx == nil {
		return nil
	}
	if session, ok := ctx.Value(contextUserSession).(UserSession); ok {
		return &session
	}
	return nil
}
