package main

import (
	"github.com/gin-gonic/gin"

	configApi "github.com/pitakill/go_api/config/api"
	"github.com/pitakill/go_api/errors"
	"github.com/pitakill/go_api/person"
	"github.com/pitakill/go_api/user"
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

	users := v1.Group("/user")
	{
		users.POST("login", user.Login)
	}

	router.NoRoute(errors.NotFound)

	router.Run("localhost:" + configApi.GetConfig())
}
