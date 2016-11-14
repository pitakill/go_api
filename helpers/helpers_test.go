package helpers

import (
	"fmt"
	"testing"
)

var validEmails = [...]string{
	"polo@vincoorbis.com",
	"mi@correo.com",
	"123@123.net",
	"a@test.com.mx",
	"email+user@dominio.com",
}

var invalidEmails = [...]string{
	"a",
	"2",
	"a@.com",
	"123asb@com",
	"132@asd@dominio.com",
}

func TestVerifyEmailFormat(t *testing.T) {
	for i := 0; i < len(validEmails); i++ {
		element := validEmails[i]
		result := VerifyEmailFormat(element)

		if result == nil {
			t.Errorf("%s is not a valid email", element)
		} else {
			fmt.Printf("Valid email: %s\n", result)
		}
	}

	for i := 0; i < len(invalidEmails); i++ {
		element := invalidEmails[i]
		result := VerifyEmailFormat(element)

		if result != nil {
			t.Errorf("%s is a valid email", result)
		} else {
			fmt.Printf("Invalid email: %s\n", element)
		}
	}
}

var validStrings = [...]string{
	"a",
	"abcedfhij",
	"This is with spaces",
	"1",
	"1234567890",
	"nil",
}

const emptyString = ""

func TestVerifyStringNull(t *testing.T) {
	for i := 0; i < len(validStrings); i++ {
		element := validStrings[i]
		result := VerifyStringNull(element)

		if result == nil {
			t.Errorf("%s is a empty string", element)
		} else {
			fmt.Printf("Is not a empty string: %s\n", result)
		}
	}

	result := VerifyStringNull(emptyString)

	if result != nil {
		t.Errorf("%s is not a empty string", result)
	} else {
		fmt.Printf("Is a empty string: '%s'\n", emptyString)
	}
}

var originalStrings = [...]string{
	"a",
	"12",
	"bla",
	"test",
}

func TestGetNewValue(t *testing.T) {
	const original = ":-)"

	for i := 0; i < len(originalStrings); i++ {
		element := originalStrings[i]
		result := GetNewValue(original, element)

		if result == original {
			t.Errorf("GetNewValue returns: '%s', but needs to return: '%s'", result, element)
		} else {
			fmt.Printf("GetNewValue('%s', '%s') returns: '%s'\n", original, element, result)
		}
	}

	const emptyString = ""
	result := GetNewValue(original, emptyString)

	if result != original {
		t.Errorf("GetNewValue returns: '%s', but needs to return: '%s'", result, original)
	} else {
		fmt.Printf("GetNewValue('%s', '%s') returns: '%s'\n", original, emptyString, result)
	}

	resultBool := GetNewValue(original, nil)

	if resultBool != original {
		t.Errorf("GetNewValue returns: '%v', but needs to return: '%s'", resultBool, original)
	} else {
		fmt.Printf("GetNewValue('%s', '%v') returns: '%s'\n", original, nil, resultBool)
	}
}
