package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	godotenv.Load(".env.local")
	BCRYPT_SALT, _ := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), BCRYPT_SALT)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
