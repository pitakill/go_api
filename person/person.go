package person

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pitakill/go_api/helpers"
	"github.com/pitakill/go_api/models"
	"github.com/pitakill/go_api/orm"
)

var db = orm.Connection()

func GetPersons(c *gin.Context) {
	var (
		person  models.Person
		persons []models.Person
	)

	rows, err := db.Query("SELECT id, first_name, last_name, email, telephone, registered FROM person;")

	if err != nil {
		fmt.Print(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&person.Id, &person.First_Name, &person.Last_Name, &person.Email, &person.Telephone, &person.Registered)
		persons = append(persons, person)

		if err != nil {
			fmt.Println(err.Error())
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": persons,
		"count":  len(persons),
	})
}

func GetPerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)

	id := c.Param("id")
	row := db.QueryRow("SELECT id, first_name, last_name, email, telephone, registered FROM person WHERE id=?;", id)
	err := row.Scan(&person.Id, &person.First_Name, &person.Last_Name, &person.Email, &person.Telephone, &person.Registered)

	if err != nil {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func PostPerson(c *gin.Context) {
	var (
		result gin.H
		status int
	)

	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	email := c.PostForm("email")
	telephone := c.PostForm("telephone")

	stmt, err := db.Prepare("INSERT INTO person (first_name, last_name, email, telephone) VALUES (?,?,?,?);")

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec(helpers.VerifyStringNull(first_name), helpers.VerifyStringNull(last_name), helpers.VerifyEmailFormat(email), helpers.VerifyStringNull(telephone))

	if err != nil {
		errorCode, message := helpers.ExtractError(err.Error())

		result = gin.H{
			"code":    errorCode,
			"message": message,
		}

		switch errorCode {
		case "1062":
			status = http.StatusUnprocessableEntity
		default:
			status = http.StatusBadRequest
		}
	} else {
		result = gin.H{
			"message": fmt.Sprintf("The person: %s %s was successfully created", first_name, last_name),
		}
		status = http.StatusCreated
	}

	c.JSON(status, result)
}

func PutPerson(c *gin.Context) {

	id := c.Param("id")

	var (
		result gin.H
		status int
	)

	// I don't know if this is the best approach
	type Data struct {
		First_Name,
		Last_Name,
		Email string
		Telephone interface{}
	}

	recorded := Data{}
	new := Data{
		c.PostForm("first_name"),
		c.PostForm("last_name"),
		c.PostForm("email"),
		c.PostForm("telephone"),
	}

	row := db.QueryRow("SELECT first_name, last_name, email, telephone FROM person WHERE id=?;", id)
	err := row.Scan(&recorded.First_Name, &recorded.Last_Name, &recorded.Email, &recorded.Telephone)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			result = gin.H{
				"code":    "404",
				"message": fmt.Sprintf("There is not person with the id: %s", id),
			}
			c.JSON(http.StatusNotFound, result)
		}
	} else {
		// There is a better way to do this
		first_name := helpers.GetNewValue(recorded.First_Name, new.First_Name)
		last_name := helpers.GetNewValue(recorded.Last_Name, new.Last_Name)
		email := helpers.GetNewValue(recorded.Email, new.Email)
		telephone := helpers.GetNewValue(recorded.Telephone, new.Telephone)

		stmt, err := db.Prepare("UPDATE person SET first_name=?, last_name=?, email=?, telephone=? WHERE id=?;")

		if err != nil {
			fmt.Println(err.Error())
		}

		_, err = stmt.Exec(first_name, last_name, helpers.VerifyEmailFormat(email.(string)), telephone, id)

		if err != nil {
			errorCode, message := helpers.ExtractError(err.Error())

			result = gin.H{
				"code":    errorCode,
				"message": message,
			}

			switch errorCode {
			case "1062":
				status = http.StatusUnprocessableEntity
			default:
				status = http.StatusBadRequest
			}
			c.JSON(status, result)
		} else {
			c.Writer.WriteHeader(http.StatusNoContent)
		}
	}
}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")

	stmt, err := db.Prepare("DELETE FROM person WHERE id=?;")

	if err != nil {
		err.Error()
	}

	_, err = stmt.Exec(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.Writer.WriteHeader(http.StatusNoContent)
	}
}
