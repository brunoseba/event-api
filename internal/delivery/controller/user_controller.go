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

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(userUseCase usecase.UserUseCase) UserController {
	return UserController{userUseCase}
}

func (uc *UserController) Login(c *gin.Context) {
	var user *entity.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := uc.userUseCase.CreateToken(user.Nickname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user id not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": token})
}

func (uc *UserController) Register(c *gin.Context) {
	var user *entity.User

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newUser, token, err := uc.userUseCase.CreateUserService(user)

	if err != nil {
		if strings.Contains(err.Error(), "user already exists") {
			c.JSON(http.StatusConflict, gin.H{"status": "409", "message": err.Error()})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"status": "502", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newUser, "token": token})
}

// obtiene user por nickname
func (uc *UserController) GetUser(c *gin.Context) {
	var nickname *entity.UserUpdate
	err := c.BindJSON(&nickname)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := uc.userUseCase.GetUser(nickname.Nickname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user id not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUserIdController(c *gin.Context) {
	userId := c.Param("id")

	if len(userId) != 24 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id doesn't have the correct length"})
	}

	user, err := uc.userUseCase.GetUserIdService(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user id not found "})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUserController(c *gin.Context) {
	userId := c.Param("id")

	if len(userId) != 24 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id doesn't have the correct length"})
	}

	var userUp *entity.UserUpdate
	if err := c.ShouldBindJSON(&userUp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "is not JSON -- " + err.Error()})
		return
	}

	updateUser, err := uc.userUseCase.UpdateUserService(userId, userUp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user id not found"})
		return
	}

	c.JSON(http.StatusOK, updateUser)

}
