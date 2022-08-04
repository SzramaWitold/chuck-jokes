package responses

import (
	"time"
)

type TokenResponse struct {
	Token      string
	TTL        *time.Time
	RefreshTTL *time.Time
}

func NewTokenResponse(token string, ttl, refreshTTL *time.Time) *TokenResponse {
	return &TokenResponse{
		Token:      token,
		TTL:        ttl,
		RefreshTTL: refreshTTL,
	}
}
