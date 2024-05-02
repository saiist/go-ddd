package handler

import (
	"go-ddd/src/application/users"
	"go-ddd/src/handler/models"
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
	result, err := h.userGetService.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.NewUserResponseModel(result)
	c.JSON(http.StatusOK, gin.H{"user": models.NewUserGetResponseModel(user)})
}

func (h *UserHandler) GetAll(c *gin.Context) {
	result, err := h.userGetService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	users := models.NewUserIndexResponseModel(result)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *UserHandler) Post(c *gin.Context) {
	var req models.UserPostRequestModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userRegisterService.Handle(req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (h *UserHandler) Put(c *gin.Context) {
	id := c.Param("id")
	var req models.UserPutRequestModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userUpdateService.Update(id, req.Name); err != nil {
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
