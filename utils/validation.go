package utils

import (
	"net/url"
	"regexp"

	"golang.org/x/exp/slices"
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

// // Cat Validation
// "name": "", // not null, minLength 1, maxLength 30
func ValidateCatName(name string) bool {
	return len(name) >= 1 && len(name) <= 30
}

// "race": "", /** not null, enum of:
//   - "Persian"
//   - "Maine Coon"
//   - "Siamese"
//   - "Ragdoll"
//   - "Bengal"
//   - "Sphynx"
//   - "British Shorthair"
//   - "Abyssinian"
//   - "Scottish Fold"
//   - "Birman" */
func ValidateCatRace(race string) bool {
	raceList := []string{"Persian", "Maine Coon", "Siamese", "Ragdoll", "Bengal", "Sphynx", "British Shorthair", "Abyssinian", "Scottish Fold", "Birman"}
	return slices.Contains(raceList, race)
}

// "sex": "", // not null, enum of: "male" / "female"
func ValidateCatSex(sex string) bool {
	sexList := []string{"male", "female"}
	return slices.Contains(sexList, sex)
}

// "ageInMonth": 1, // not null, min: 1, max: 120082
func ValidateCatAgeInMonth(age int32) bool {
	return age >= 1 && age <= 120082
}

// "description":"" // not null, minLength 1, maxLength 200
func ValidateCatDescription(description string) bool {
	return len(description) >= 1 && len(description) <= 200
}

// "imageUrls":[ // not null, minItems: 1, items: not null, should be url
//
//	"","",""
//
// ]
func ValidateCatImageUrls(images []string) bool {
	if len(images) == 0 {
		return false
	}
	for i := 0; i < len(images); i++ {
		_, err := url.ParseRequestURI(images[i])
		if err != nil {
			return false
		}
	}
	return true
}
