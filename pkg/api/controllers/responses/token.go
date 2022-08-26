package responses

import (
	"time"
)

type Token struct {
	Token      string
	TTL        *time.Time
	RefreshTTL *time.Time
}

func (r *Response) NewToken(token string, ttl, refreshTTL *time.Time) Token {
	return Token{
		Token:      token,
		TTL:        ttl,
		RefreshTTL: refreshTTL,
	}
}
