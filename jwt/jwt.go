package jwt

import (
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var hmacSampleSecret []byte

func CreateToken(id int, email string) string {
	if keyData, e := ioutil.ReadFile("jwtKey"); e == nil {
		hmacSampleSecret = keyData
	}

	now := time.Now().UTC().Unix()
	const validUntil int64 = 60 * 60 // 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   now + validUntil,
		"id":    id,
		"nbf":   now,
	})

	tokenString, _ := token.SignedString(hmacSampleSecret)

	return tokenString
}
