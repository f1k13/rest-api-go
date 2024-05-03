package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte("ARCHLINUX")

func GenerateToken(userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenStr, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ValidateToken(token string) (string, error) {
	if token == "" {
		return "", jwt.ErrSignatureInvalid
	}
	tokenVal, err := jwt.Parse(token, func(tokenVal *jwt.Token) (interface{}, error) {
		return []byte("ARCHLINUX"), nil
	})
	if err != nil {
		return "", err
	}
	if !tokenVal.Valid {
		return "", jwt.ErrSignatureInvalid
	}

	claims, ok := tokenVal.Claims.(jwt.MapClaims)
	if !ok {
		return "", jwt.ErrInvalidKeyType
	}
	userID, ok := claims["id"].(string)
	if !ok {
		return "", jwt.ErrInvalidKeyType
	}
	return userID, nil
}
