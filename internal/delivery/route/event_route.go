package route

import (
	"github.com/brunoseba/event-api/internal/delivery/controller"
	"github.com/brunoseba/event-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

type EventRouteConstroller struct {
	eventController controller.EventController
}

func NewEventRouteController(eventController controller.EventController) EventRouteConstroller {
	return EventRouteConstroller{eventController}
}

func (er *EventRouteConstroller) EventRoute(rg *gin.RouterGroup) {
	routes := rg.Group("/v1")

	routes.GET("/event/:id", middleware.UserTypeMiddleware(), er.eventController.GetEvent)
	routes.GET("/events/", middleware.AuthUser(), er.eventController.GetAllEventController)

	routes.GET("/eventfilt", middleware.UserTypeMiddleware(), er.eventController.GetEventController)
	routes.POST("/event", middleware.AuthUser(), er.eventController.CreateEventController)
	routes.PUT("/event/update/:id", middleware.AuthUser(), er.eventController.UpdateEventController)
	routes.POST("/event/inscription", er.eventController.InscriptionUserEvent)
	routes.GET("/eventsuser/:id", er.eventController.GetListEventsUserController)
	routes.DELETE("/event/:id", er.eventController.DeleteEventController)

}
