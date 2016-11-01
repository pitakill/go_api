package main

import (
	"github.com/gin-gonic/gin"

	"github.com/pitakill/go_api/person"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	v1 := api.Group("/v1")
	{
		v1.GET("persons", person.GetPersons)
		v1.GET("person/:id", person.GetPerson)
		v1.POST("person", person.PostPerson)
		v1.PUT("person/:id", person.PutPerson)
		v1.DELETE("person/:id", person.DeletePerson)
	}

	router.Run(":3000")
}
