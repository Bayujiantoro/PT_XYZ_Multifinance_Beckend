package regexp

import "regexp"

func EmailValidation(email string) bool {
	validation := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	result := validation.MatchString(email)

	return result
}