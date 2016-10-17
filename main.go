package main

import (
	"github.com/gin-gonic/gin"

	"github.com/pitakill/go_api/person"
)

func main() {
	router := gin.Default()

	router.GET("persons", person.GetPersons)
	router.GET("person/:id", person.GetPerson)
	router.POST("person", person.PostPerson)
	router.PUT("person/:id", person.PutPerson)
	router.DELETE("person/:id", person.DeletePerson)

	router.Run(":3000")
}
