package CH8

import (
	"unicode"
)

func isValidPassword(password string) bool {
	if !(len(password) > 5 && len(password) < 12) {
		return false
	}
	hasDigit := false
	hasUppercase := false
	for _, c := range password {
		if unicode.IsDigit(c) {
			hasDigit = true
		}
		if unicode.IsUpper(c) {
			hasUppercase = true
		}
		if hasDigit && hasUppercase {
			break
		}
	}
	return hasDigit && hasUppercase
}
