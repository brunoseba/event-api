package route

import (
	"github.com/brunoseba/event-api/internal/delivery/controller"
	"github.com/brunoseba/event-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controller.UserController
}

func NewUserRouteController(userController controller.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (ur *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("/v1")

	router.POST("/login", ur.userController.Login)

	router.GET("/user", ur.userController.GetUser)
	router.GET("/user/:id", ur.userController.GetUserIdController)
	router.POST("/register", ur.userController.Register)
	router.PUT("/user/update/:id", middleware.AuthUser(), ur.userController.UpdateUserController)
}
