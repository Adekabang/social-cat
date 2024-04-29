package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GetToken(username, jwtSecret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username

	return token.SignedString([]byte(jwtSecret))
}

func GetJWTSecret() (string, error) {
	godotenv.Load(".env.local")

	jwtSecret, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		return "", errors.New("environment variable JWT_SECRET not set")
	}

	return jwtSecret, nil
}
