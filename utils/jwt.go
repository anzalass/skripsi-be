package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTInterface interface {
	GenerateJWT(name, email, role string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) JWTInterface {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) GenerateJWT(name string, email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"name":  name,
		"email": email,
		"role":  role,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, _ := token.SignedString([]byte(j.Secret))
	return accessToken, nil
}

func (j *JWT) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
