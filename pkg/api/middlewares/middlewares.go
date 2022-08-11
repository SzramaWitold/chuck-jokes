package middlewares

import (
	"chuck-jokes/pkg/token"
)

type Middleware struct {
	Auth *AuthenticationMiddleware
}

func NewMiddleware(jwt *token.Handler) *Middleware {
	return &Middleware{Auth: NewAuthenticationMiddleware(jwt)}
}
