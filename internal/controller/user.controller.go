package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/huynhbaoking112/System_design_Golang/internal/service"
	"github.com/huynhbaoking112/System_design_Golang/package/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserRepoService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	fmt.Println("MyHandler")
	response.SuccessResponse(c, 20001, []string{"king", "huynh"})
}
