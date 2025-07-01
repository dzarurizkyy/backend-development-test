package utils

import (
	"api-test/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id int) (string, error) {
	jwtSecret := []uint8(config.AppConfig.JWT.Secret)

	claims := jwt.MapClaims{
		"admin_id": id,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
