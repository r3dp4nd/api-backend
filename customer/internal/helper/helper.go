package helper

import (
	"regexp"
	"time"
)

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(EmailPattern)
	return emailRegex.MatchString(email)
}

func IsValidAge(birthday time.Time, minYear int) bool {
	age := time.Now().Year() - birthday.Year()

	if age <= minYear {
		return true
	}

	return false
}
