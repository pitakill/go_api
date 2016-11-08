package jwt

import (
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/pitakill/go_api/models"
)

var hmacSampleSecret []byte

func CreateToken(user models.User) string {
	if keyData, e := ioutil.ReadFile("jwtKey"); e == nil {
		hmacSampleSecret = keyData
	}

	now := time.Now().UTC().Unix()
	const validUntil int64 = 60 * 60 // 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  now + validUntil,
		"nbf":  now,
		"user": user,
	})

	tokenString, _ := token.SignedString(hmacSampleSecret)

	return tokenString
}
