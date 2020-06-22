package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jayjzheng/amazon/auth/internal/domain"
)

const (
	defaultDuration = 5 * time.Minute
)

type JWT struct {
	TokenDuration time.Duration
	Secret        []byte

	signMethod jwt.SigningMethod
}

func NewJWT(opts ...func(*JWT)) *JWT {
	var j JWT

	for _, opt := range opts {
		opt(&j)
	}

	if j.TokenDuration == 0 {
		j.TokenDuration = defaultDuration
	}

	j.signMethod = jwt.SigningMethodHS256

	return &j
}

func WithSecret(s []byte) func(*JWT) {
	return func(j *JWT) {
		j.Secret = s
	}
}

func WithTokenDuration(d time.Duration) func(*JWT) {
	return func(j *JWT) {
		j.TokenDuration = d
	}
}

type claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

func (j *JWT) GenerateToken(u *domain.User) (string, error) {
	expiry := time.Now().Add(j.TokenDuration)

	cc := &claims{
		Login: u.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
		},
	}

	token := jwt.NewWithClaims(j.signMethod, cc)
	str, err := token.SignedString(j.Secret)
	if err != nil {
		return "", fmt.Errorf("SignedString: %w", err)
	}

	return str, nil
}
