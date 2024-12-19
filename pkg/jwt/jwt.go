package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username string, seckerKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"exp":      time.Now().Add(1 * time.Minute).Unix(),
		},
	)

	key := []byte(seckerKey)

	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateToken(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", err
	}

	id := int64(claims["id"].(float64))
	username := claims["username"].(string)

	return id, username, nil
}

func ValidateTokenWithoutExpiry(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, jwt.WithoutClaimsValidation())

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", err
	}

	id := int64(claims["id"].(float64))
	username := claims["username"].(string)

	return id, username, nil
}
