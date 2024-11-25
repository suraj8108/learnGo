package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suraj8108/learngo/Gorestproject/db"
	"github.com/suraj8108/learngo/Gorestproject/routes"
)

func main() {
	DB := db.InitDB()
	server := gin.Default() // Creates the logger and all other basic functions required to start the application

	eventHandler := &routes.EventHandler{
		DB: DB,
	}

	userHandler := &routes.UserHandler{
		DB: DB,
	}

	eventRegistrationHandler := &routes.EventRegisterHandler{
		DB: DB,
	}
	routes.RegisterRoutes(server, eventHandler, eventRegistrationHandler, userHandler)

	server.Run(":8080")

}
