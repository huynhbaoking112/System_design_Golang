package user

import (
	"github.com/gin-gonic/gin"
	"github.com/huynhbaoking112/System_design_Golang/internal/wire"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := wire.InitUserRouterHandler()

	// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	// useRouterPrivate.Use(limiter())
	// useRouterPrivate.Use(Authen())
	// useRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.GET("/otp")
	}

}
