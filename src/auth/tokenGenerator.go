package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/taral14/filmun/src/entity"
)

type Token struct {
	UserId int
	jwt.StandardClaims
}

type JwtTokenGenerator struct {
	signingKey     []byte
	expireDuration time.Duration
}

func NewJwtTokenGenerator(signingKey []byte, tokenTTLSeconds time.Duration) *JwtTokenGenerator {
	return &JwtTokenGenerator{
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

func (p *JwtTokenGenerator) CreateToken(user *entity.User) (string, error) {
	claims := Token{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(p.expireDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(p.signingKey)
}

func (p *JwtTokenGenerator) ParseToken(tokenString string) (*Token, error) {
	claims := &Token{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return p.signingKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Cant parse token: %w", err)
	}
	if !token.Valid {
		return nil, errors.New("JwtTokenGenerator: Token is not valid")
	}
	return claims, nil
}
