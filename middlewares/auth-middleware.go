package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
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

func ValidateToken(token string) (int, error) {
	if token == "" {
		return 0, jwt.ErrSignatureInvalid
	}
	tokenVal, err := jwt.Parse(token, func(tokenVal *jwt.Token) (interface{}, error) {
		return []byte("ARCHLINUX"), nil
	})
	if err != nil {
		logrus.Error(err.Error())
		return 0, err
	}
	if !tokenVal.Valid {
		logrus.Error("Инвалид токен")
		return 0, jwt.ErrSignatureInvalid
	}

	claims, ok := tokenVal.Claims.(jwt.MapClaims)
	if !ok {
		logrus.Error("Ошибка claims")
		return 0, jwt.ErrInvalidKeyType
	}
	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		logrus.Error("Ошибка userID")
		return 0, jwt.ErrInvalidKeyType
	}
	userID := int(userIDFloat)
	return userID, nil
}
