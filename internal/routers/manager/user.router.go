package manager

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register")
		// userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := Router.Group("/admin/user")
	// useRouterPrivate.Use(limiter())
	// useRouterPrivate.Use(Authen())
	// useRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user")
	}

}
