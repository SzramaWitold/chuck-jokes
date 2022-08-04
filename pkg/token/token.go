package token

import (
	modelsGorm "chuck-jokes/pkg/database/models/gorm"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"strconv"
	"time"
)

type Validator struct {
	secret []byte
}

func NewValidator() *Validator {
	return &Validator{
		secret: []byte(os.Getenv("SECRET")),
	}
}

func (v *Validator) CreateToken(user *modelsGorm.User) (string, *time.Time, *time.Time) {
	ttlDuration, refreshTTLDuration := setTTLAndRefresh()

	ttl := time.Now().Add(time.Duration(ttlDuration) * time.Minute)
	refreshTtl := time.Now().Add(time.Duration(refreshTTLDuration) * time.Minute)
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID":      user.ID,
			"ttl":         strconv.Itoa(int(ttl.Unix())),
			"refresh_ttl": strconv.Itoa(int(refreshTtl.Unix())),
		},
	)

	tokenString, stringTokenError := token.SignedString(v.secret)

	if stringTokenError != nil {
		log.Println(stringTokenError)
	}

	return tokenString, &ttl, &refreshTtl
}

func (v *Validator) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return v.secret, nil
	})

	if tokenErr != nil {
		log.Println(tokenErr)
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

func setTTLAndRefresh() (int, int) {
	ttlDuration := 5
	refreshTTLDuration := 10
	if os.Getenv("TTL") != "" {
		newTTL, ttlErr := strconv.Atoi(os.Getenv("TTL"))
		if ttlErr != nil {
			log.Println(ttlErr)
		} else {
			ttlDuration = newTTL
		}
	}

	if os.Getenv("TTL") != "" {
		newRefreshTTL, ttlErr := strconv.Atoi(os.Getenv("REFRESH_TTL"))
		if ttlErr != nil {
			log.Println(ttlErr)
		} else {
			refreshTTLDuration = newRefreshTTL
		}
	}

	return ttlDuration, refreshTTLDuration
}
