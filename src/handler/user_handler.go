package handler

import (
	"go-ddd/src/application/dtos"
	"go-ddd/src/application/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userGetService      *users.UserGetService
	userRegisterService *users.UserRegisterService
	userUpdateService   *users.UserUpdateService
	userDeleteService   *users.UserDeleteService
}

func NewUserHandler(
	userGetService *users.UserGetService,
	userRegisterService *users.UserRegisterService,
	userUpdateService *users.UserUpdateService,
	userDeleteService *users.UserDeleteService,
) *UserHandler {
	return &UserHandler{
		userGetService:      userGetService,
		userRegisterService: userRegisterService,
		userUpdateService:   userUpdateService,
		userDeleteService:   userDeleteService,
	}
}

func (h *UserHandler) Get(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userGetService.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) Post(c *gin.Context) {
	var user dtos.UserData
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userRegisterService.Handle(user.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (h *UserHandler) Put(c *gin.Context) {
	id := c.Param("id")
	var user dtos.UserData
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userUpdateService.Update(id, user.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.userDeleteService.Handle(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
