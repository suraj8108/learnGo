package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suraj8108/learngo/Gorestproject/middleware"
)

func RegisterRoutes(server *gin.Engine, eventHandler *EventHandler, eventRegistrationHandler *EventRegisterHandler, userHandler *UserHandler) {

	server.GET("/events", eventHandler.GetEvents)
	server.GET("/events/:id", eventHandler.GetEventById) // Get Event By ID

	//Grouped the routes
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate) // Adding middleware in one go to multiple routes
	authenticated.POST("/events", eventHandler.CreateEvent)
	authenticated.PUT("/events/:id", eventHandler.UpdateEvent)
	authenticated.DELETE("/events/:id", eventHandler.DeleteEvent)

	authenticated.POST("/events/:id/register", eventRegistrationHandler.registerForEvent)
	authenticated.DELETE("/events/:id/register", eventRegistrationHandler.cancellinRegistration)
	// server.POST("/events", middleware.Authenticate, eventHandler.CreateEvent)
	// server.PUT("/events/:id", eventHandler.UpdateEvent)
	// server.DELETE("/events/:id", eventHandler.DeleteEvent)

	server.POST("/signup", userHandler.signUp)
	server.POST("/login", userHandler.login)
}
