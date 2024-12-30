package CH8

import (
	"errors"
	"testing"
)

func TestTagMessages(t *testing.T) {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	test12("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	test12("Error on mail server", mailErrors, commaDelimit)
}
