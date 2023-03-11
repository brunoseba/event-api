package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/brunoseba/event-api/internal/entity"
	"github.com/brunoseba/event-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	eventUseCase usecase.EventUseCase
}

func NewEventController(eventUseCase usecase.EventUseCase) EventController {
	return EventController{eventUseCase}
}

func (ec *EventController) CreateEventController(c *gin.Context) {
	var event *entity.Event
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &event)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newEvent, err := ec.eventUseCase.CreateEventService(event)

	if err != nil {
		if strings.Contains(err.Error(), "user already exists") {
			c.JSON(http.StatusConflict, gin.H{"status": "409", "message": err.Error()})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"status": "502", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newEvent})

}

func (ec *EventController) GetEvent(c *gin.Context) {
	eventId := c.Param("id")
	usertype := c.GetBool("enabled")

	event, err := ec.eventUseCase.GetEvent(eventId, usertype)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get event id not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (ec *EventController) GetAllEventController(c *gin.Context) {

	events, err := ec.eventUseCase.GetAllEventService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get all events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}

func (ec *EventController) GetEventController(c *gin.Context) {
	var filt *entity.EventFilt
	err := c.BindJSON(&filt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var (
		dataIn   = filt.DateIn
		dataOut  = filt.DateOut
		title    = filt.Title
		state    = filt.State
		usertype = c.GetBool("enabled")
	)

	events, err := ec.eventUseCase.GetEventFiltDateService(usertype, dataIn, dataOut, title, state)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get events"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}

func (ec *EventController) UpdateEventController(c *gin.Context) {
	eventId := c.Param("id")
	if len(eventId) != 24 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id doesn't have the correct length"})
	}

	var eventUp *entity.EventUpdate
	if err := c.ShouldBindJSON(&eventUp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "is not JSON -- " + err.Error()})
		return
	}

	eventUpdate, err := ec.eventUseCase.UpdateEvent(eventId, eventUp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update event"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event update", "data": eventUpdate})
}

func (ec *EventController) InscriptionUserEvent(c *gin.Context) {
	var eventuser *entity.UserEvent
	err := c.BindJSON(&eventuser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ec.eventUseCase.AddEventToUser(eventuser.UserId, eventuser.EventId)

	c.JSON(http.StatusOK, gin.H{})
}

func (ec *EventController) GetListEventsUserController(c *gin.Context) {
	userId := c.Param("id")

	if len(userId) != 24 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id doesn't have the correct length"})
	}
	events, err := ec.eventUseCase.GetListEventUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get events"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": events})
}

func (ec *EventController) DeleteEventController(c *gin.Context) {
	eventid := c.Param("id")
	err := ec.eventUseCase.DeleteEventById(eventid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to delete event"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
