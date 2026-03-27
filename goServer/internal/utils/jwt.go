package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accessSecret = []byte("access_secret")
var refreshSecret = []byte("refresh_secret")

func GenerateTokens(email string) (string, string, error) {

	// Access Token (1 day)
	accessClaims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString(accessSecret)
	if err != nil {
		return "", "", err
	}

	// Refresh Token (7 days)
	refreshClaims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshStr, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		return "", "", err
	}

	return accessStr, refreshStr, nil
}
