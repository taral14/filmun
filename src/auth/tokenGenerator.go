package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/taral14/filmun/src/entity"
)

type JwtClaims struct {
	jwt.StandardClaims
	User *entity.User `json:"user"`
}

type JwtTokenGenerator struct {
	signingKey     []byte
	expireDuration time.Duration
}

func NewJwtTokenGenerator(
	signingKey []byte,
	tokenTTLSeconds time.Duration) *JwtTokenGenerator {
	return &JwtTokenGenerator{
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

func (p *JwtTokenGenerator) CreateToken(user *entity.User) (string, error) {
	claims := JwtClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(p.expireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(p.signingKey)
}
