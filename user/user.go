package user

import (
	"crypto/sha512"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pitakill/go_api/config/message"
	"github.com/pitakill/go_api/jwt"
	"github.com/pitakill/go_api/models"
	"github.com/pitakill/go_api/orm"
)

var db = orm.Connection()

func sha2Password(password string) (hash string) {
	inputPassword := sha512.Sum512([]byte(password))
	passwordEncoded := inputPassword[:]

	hash = hex.EncodeToString(passwordEncoded)
	return
}

func Login(c *gin.Context) {
	var (
		user   models.User
		result gin.H
		status int
	)

	email, password := c.PostForm("email"), c.PostForm("password")

	passwordHashed := sha2Password(password)

	row := db.QueryRow("SELECT id, first_name, last_name, email, username, type, twitter from user WHERE email=? AND password=?;", email, passwordHashed)
	err := row.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Username, &user.Type, &user.Twitter)

	configMessage := message.GetMessagesUser()

	if err != nil {
		status = http.StatusNotFound
		result = gin.H{
			"code":    configMessage.User.Login.Error.Code,
			"message": configMessage.User.Login.Error.Message,
		}
	} else {
		status = http.StatusOK
		result = gin.H{
			"token": jwt.CreateToken(user),
		}
	}

	c.JSON(status, result)
}

//func PostUser(c *gin.Context) {

//}
