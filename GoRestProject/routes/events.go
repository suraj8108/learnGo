package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suraj8108/learngo/Gorestproject/models"
)

type EventHandler struct {
	DB *sql.DB
}

func (e EventHandler) GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents(e.DB)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func (e EventHandler) GetEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventId from params"})
		return
	}
	event, err := models.GetEventById(eventId, e.DB)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"eventDetails": event})
}

func (e EventHandler) CreateEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save(e.DB)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func (e EventHandler) UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventId from params"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId, e.DB)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authrorized to update event"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data"})
		return
	}

	updatedEvent.ID = eventId
	err = models.UpdateEvent(updatedEvent, e.DB)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not update the event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event updated Successfully", "event": updatedEvent})

}

func (e EventHandler) DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventId from params"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId, e.DB)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could find event with the given eventID"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authrorized to update event"})
		return
	}

	err = models.DeleteEvent(eventId, e.DB)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not dekete an event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event deleted Successfully"})

}
