package middlewares

import (
	"chuck-jokes/pkg/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

type AuthenticationMiddleware struct {
	JWT *token.Handler
}

func NewAuthenticationMiddleware(JWT *token.Handler) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{JWT: JWT}
}

func (mid *AuthenticationMiddleware) Auth(c *gin.Context) {
	const BearerSchema = "Bearer "
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]

	t, tokenErr := mid.JWT.ValidateToken(tokenString)
	if tokenErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, tokenErr.Error())
		return
	} else if !t.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthenticated")
		return
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		c.AddParam("userID", fmt.Sprintf("%v", claims["userID"]))
	}

	c.Next()
}
