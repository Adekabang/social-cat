package utils

import (
	"regexp"
)

// ValidateEmail checks if the email is in the correct format
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePassword checks if the password meets length requirements
func ValidatePassword(password string) bool {
	return len(password) >= 5 && len(password) <= 15
}

// ValidateName checks if the name meets length requirements
func ValidateName(name string) bool {
	return len(name) >= 5 && len(name) <= 50
}
