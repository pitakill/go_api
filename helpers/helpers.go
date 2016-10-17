package helpers

import (
	"regexp"
	"strings"
)

func VerifyEmailFormat(s string) interface{} {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if Re.MatchString(s) {
		return s
	} else {
		return nil
	}
}

func VerifyStringNull(s string) interface{} {
	if s != "" {
		return s
	} else {
		return nil
	}
}

func GetNewValue(a, b interface{}) interface{} {
	if b == "" || b == nil {
		return a
	} else {
		return b
	}
}

func ExtractError(s string) (code, field string) {
	stringSplited := strings.Split(s, " ")
	errorCode := strings.TrimSuffix(stringSplited[1], ":")
	errorMessage := strings.Split(s, ":")
	return errorCode, strings.TrimPrefix(errorMessage[1], " ")
}
