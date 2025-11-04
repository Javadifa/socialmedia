package email

import "fmt"

func isEmailValid(email string) (bool, error) {
	//TODO - technical debt - is EMAIL valid
	return false, fmt.Errorf("email is not valid")
}

func isEmailUnique(email string) (bool, error) {
	//TODO - technical debt - is EMAIL UNIQUE
	return false, fmt.Errorf("email is not unique")
}
