package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GetSecretKey returns a key for creating and validating JWT tokens.
// If secret key doesn't exist, it will be generated.
func GetSecretKey() string {
	return "SOME_JWT_SECRET_KEY"
}

// GenerateToken generates JWT token via nickname and current server time.
func GenerateToken(nickname string) string {
	data := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		Issuer:    nickname,
		IssuedAt:  time.Now().Unix(),
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	token, err := rawToken.SignedString([]byte(GetSecretKey()))
	if err != nil {
		panic(err)
	}
	return token
}

// ValidateToken validates provided token.
func ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validating signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(GetSecretKey()), nil
	})
}

// ParseToken parses provided token and returns its claims.
func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetSecretKey()), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
