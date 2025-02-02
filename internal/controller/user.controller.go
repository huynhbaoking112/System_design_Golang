package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhbaoking112/System_design_Golang/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserRepo(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUserService(),
		"uid":     uid,
		"users": []string{
			"cr7",
			"m10",
			"king",
		},
	})
}
