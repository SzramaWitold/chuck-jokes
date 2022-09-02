package token

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"chuck-jokes/models"
	"github.com/golang-jwt/jwt/v4"
)

var user = models.User{
	ID: 5,
}

var (
	secret     = []byte("test")
	now        = time.Now()
	ttl        = now.Add(time.Duration(5) * time.Minute)
	refreshTtl = now.Add(time.Duration(10) * time.Minute)
	token      = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserID":      user.ID,
		"ttl":         strconv.Itoa(int(ttl.Unix())),
		"refresh_ttl": strconv.Itoa(int(refreshTtl.Unix())),
	})
)

func TestHandler_CreateToken(t *testing.T) {
	tokenString, _ := token.SignedString(secret)
	type fields struct {
		secret     []byte
		ttl        int
		refreshTtl int
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  *time.Time
		want2  *time.Time
	}{
		{
			name: "Base test",
			fields: fields{
				secret:     secret,
				ttl:        5,
				refreshTtl: 10,
			},
			args: args{
				user: &user,
			},
			want:  tokenString,
			want1: &ttl,
			want2: &refreshTtl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				secret:     tt.fields.secret,
				ttl:        tt.fields.ttl,
				refreshTtl: tt.fields.refreshTtl,
			}
			got, got1, got2 := h.CreateToken(tt.args.user)
			if got != tt.want {
				t.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			}
			parsedGot1 := got1.Format(time.UnixDate)
			parsedWant1 := tt.want1.Format(time.UnixDate)
			parsedGot2 := got2.Format(time.UnixDate)
			parsedWant2 := tt.want2.Format(time.UnixDate)

			if parsedGot1 != parsedWant1 {
				t.Errorf("CreateToken() got1 = %v, want %v", got1, tt.want1)
			}
			if parsedGot2 != parsedWant2 {
				t.Errorf("CreateToken() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestHandler_ValidateToken(t *testing.T) {
	tokenString, _ := token.SignedString(secret)
	type fields struct {
		secret     []byte
		ttl        int
		refreshTtl int
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *jwt.Token
		wantErr bool
	}{
		{
			name: "Base test",
			fields: fields{
				secret:     secret,
				ttl:        5,
				refreshTtl: 10,
			},
			args:    args{tokenString: tokenString},
			want:    token,
			wantErr: false,
		},
		{
			name: "Test Error",
			fields: fields{
				secret:     secret,
				ttl:        5,
				refreshTtl: 10,
			},
			args:    args{tokenString: "wrongToken"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				secret:     tt.fields.secret,
				ttl:        tt.fields.ttl,
				refreshTtl: tt.fields.refreshTtl,
			}
			got, err := h.ValidateToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil && tt.want == nil {
				return
			}
			if got.Claims.Valid() != tt.want.Claims.Valid() {
				t.Errorf("ValidateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHandler(t *testing.T) {
	type args struct {
		secret     string
		ttl        int
		refreshTTL int
	}
	tests := []struct {
		name string
		args args
		want TokenHandler
	}{
		{
			name: "Base test",
			args: args{
				secret:     "test",
				ttl:        5,
				refreshTTL: 10,
			},
			want: Handler{
				secret:     []byte("test"),
				ttl:        5,
				refreshTtl: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.secret, tt.args.ttl, tt.args.refreshTTL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
