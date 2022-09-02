package middlewares

import (
	"chuck-jokes/pkg/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

type AuthenticationMiddleware struct {
	JWT *token.TokenHandler
}

func NewAuthenticationMiddleware(JWT *token.TokenHandler) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{JWT: JWT}
}

func (mid *AuthenticationMiddleware) Auth(c *gin.Context) {
	const BearerSchema = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthenticated")
		return
	}
	tokenString := authHeader[len(BearerSchema):]

	baseJwt := *mid.JWT
	t, tokenErr := baseJwt.ValidateToken(tokenString)
	if tokenErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, tokenErr.Error())
		return
	} else if !t.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthenticated")
		return
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		c.AddParam("UserID", fmt.Sprintf("%v", claims["UserID"]))
	}

	c.Next()
}
