package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suraj8108/learngo/Gorestproject/models"
)

type EventRegisterHandler struct {
	DB *sql.DB
}

func (r EventRegisterHandler) registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can not parse event Id"})
		return
	}

	_, err = models.GetEventById(eventId, r.DB)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect eventId passed", "errorMsg": err.Error()})
		return
	}

	registerDetails := &models.RegisterEvent{
		EventId: eventId,
		UserId:  userId,
	}
	err = registerDetails.RegisterUserForEvent(r.DB)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to register a user for event", "errorMsg": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User registered for an event"})
}

func (r EventRegisterHandler) cancellinRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can not parse event Id"})
		return
	}

	_, err = models.GetEventById(eventId, r.DB)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect eventId passed", "errorMsg": err.Error()})
		return
	}

	registerDetails := &models.RegisterEvent{
		EventId: eventId,
		UserId:  userId,
	}

	//Check the registration details
	registeredEventDetails, err := registerDetails.GetRegistrationOfUserForEvent(r.DB)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No such registration found", "errorMsg": err.Error()})
		return
	}

	if userId != registeredEventDetails.UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to cancel the event", "errorMsg": err.Error()})
		return
	}

	err = registerDetails.CancelRegistration(r.DB)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to cancel user registration", "errorMsg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User registration cancelled successfully"})
}
