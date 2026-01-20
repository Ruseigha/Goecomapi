package jwtpkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret     string
	expiration time.Duration
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewJWT(secret string, expiration time.Duration) *JWT {
	return &JWT{secret: secret, expiration: expiration}
}

func (j *JWT) GenerateToken(c CustomClaims) (string, error) {
	c.ExpiresAt = jwt.NewNumericDate(time.Now().Add(j.expiration))
	c.IssuedAt = jwt.NewNumericDate(time.Now())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(j.secret))
}

func (j *JWT) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
