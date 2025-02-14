package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/huynhbaoking112/System_design_Golang/internal/service"
	"github.com/huynhbaoking112/System_design_Golang/package/response"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}
