package utils

import (
	"encoding/base64"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type Token string

type JwtClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func EncodeBasicAuth(username, password string) string {
	token := base64.StdEncoding.EncodeToString([]byte(strings.Join([]string{username, password}, ":")))

	return token
}

func DecodeBasicAuth(token string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
