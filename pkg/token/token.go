package token

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"chuck-jokes/models"

	"github.com/golang-jwt/jwt/v4"
)

type TokenHandler interface {
	CreateToken(user *models.User) (string, *time.Time, *time.Time)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type Handler struct {
	secret     []byte
	ttl        int
	refreshTtl int
}

func NewHandler(secret string, ttl, refreshTTL int) TokenHandler {
	return Handler{
		secret:     []byte(secret),
		ttl:        ttl,
		refreshTtl: refreshTTL,
	}
}

func (h Handler) CreateToken(user *models.User) (string, *time.Time, *time.Time) {
	ttl := time.Now().Add(time.Duration(h.ttl) * time.Minute)
	refreshTtl := time.Now().Add(time.Duration(h.refreshTtl) * time.Minute)
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"UserID":      user.ID,
			"ttl":         strconv.Itoa(int(ttl.Unix())),
			"refresh_ttl": strconv.Itoa(int(refreshTtl.Unix())),
		},
	)

	tokenString, stringTokenError := token.SignedString(h.secret)

	if stringTokenError != nil {
		log.Println(stringTokenError)
	}

	return tokenString, &ttl, &refreshTtl
}

func (h Handler) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return h.secret, nil
	})

	if tokenErr != nil {
		log.Println(tokenErr)
		log.Println("Token is not nil", tokenErr != nil)
		return nil, tokenErr
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		now := time.Now().Unix()
		strTTL := fmt.Sprintf("%v", claims["ttl"])
		ttl, ttlErr := strconv.Atoi(strTTL)
		if ttlErr != nil {
			return nil, ttlErr
		}

		if now > int64(ttl) {
			return nil, fmt.Errorf("token expired")
		}
	}

	return token, nil
}
