package middlewares

import (
	"chuck-jokes/di"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func (m *Middleware) Auth(c *gin.Context) {
	const BearerSchema = "Bearer "
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BearerSchema):]
	validator := di.JWT()
	t, tokenErr := validator.ValidateToken(tokenString)
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
