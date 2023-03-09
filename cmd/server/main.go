package main

import (
	"context"
	"log"

	"github.com/brunoseba/event-api/internal/config"
	"github.com/brunoseba/event-api/internal/delivery/controller"
	"github.com/brunoseba/event-api/internal/delivery/route"
	"github.com/brunoseba/event-api/internal/repository"
	"github.com/brunoseba/event-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server *gin.Engine
	ctx    context.Context

	//User Services, controllers, routes
	userCollection  *mongo.Collection
	eventCollection *mongo.Collection

	userService         usecase.UserUseCase
	UserController      controller.UserController
	UserRouteController route.UserRouteController
	UserRepository      repository.UserRepositoryImp

	eventService         usecase.EventUseCase
	EventController      controller.EventController
	EventRouteController route.EventRouteConstroller
	EventRepository      repository.EventRepositoryImp
)

func init() {

	ctx = context.TODO()

	// connect with database --------------------------------------------------------
	mongoclient, _ := config.ConnectToMongoDB(ctx)

	//Collections
	userCollection = mongoclient.Database("eventmanagement").Collection("users")
	eventCollection = mongoclient.Database("eventmanagement").Collection("event")

	//Route, Controller and Service --------------------------------------------------
	UserRepository = repository.NewUserRepository(userCollection, ctx)
	userService = usecase.NewUserUseCase(UserRepository)
	UserController = controller.NewUserController(userService)
	UserRouteController = route.NewUserRouteController(UserController)

	EventRepository = repository.NewEventReposiroty(eventCollection, ctx)
	eventService = usecase.NewEventUseCase(EventRepository)
	EventController = controller.NewEventController(eventService)
	EventRouteController = route.NewEventRouteController(EventController)

	// -------------------------------------------------------------------------------

	server = gin.Default()

}

func main() {

	// Routes
	router := server.Group("/api")
	UserRouteController.UserRoute(router)
	EventRouteController.EventRoute(router)

	log.Fatal(server.Run(":8090"))
}
