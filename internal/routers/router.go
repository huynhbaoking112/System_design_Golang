package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/huynhbaoking112/System_design_Golang/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", controller.GetPongController().Pong)
		v1.PUT("/ping", controller.NewUserController().GetUserById)
		// v1.POST("/ping", controller.Pong)
		// v1.PATCH("/ping", controller.Pong)
		// v1.HEAD("/ping", controller.Pong)
		// v1.OPTIONS("/ping", controller.Pong)
	}

	return r

}
