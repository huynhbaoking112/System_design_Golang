package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/huynhbaoking112/System_design_Golang/global"
	"github.com/huynhbaoking112/System_design_Golang/internal/routers"
)

func InitRouter() *gin.Engine {
	// r := gin.Default()
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	// r.Use() //logging
	// r.Use() //cross
	// r.Use() //limiter global

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/checkStatus") //tracking monitor
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserRouter(MainGroup)
	}

	return r

}
