package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/huynhbaoking112/System_design_Golang/internal/controller"
	"github.com/huynhbaoking112/System_design_Golang/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}
func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}
func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	// use the middleware
	r.Use(AA(), middlewares.AuthenMiddleware(), BB(), CC)

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", controller.GetPongController().Pong)
		v1.GET("/user", controller.NewUserController().GetUserById)
		// v1.POST("/ping", controller.Pong)
		// v1.PATCH("/ping", controller.Pong)
		// v1.HEAD("/ping", controller.Pong)
		// v1.OPTIONS("/ping", controller.Pong)
	}

	return r

}
