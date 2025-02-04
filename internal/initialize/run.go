package initialize

import (
	"github.com/huynhbaoking112/System_design_Golang/global"
	"go.uber.org/zap"
)

func Run() {
	InitLoadConfig()
	// fmt.Println("Loading configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config Log Ok!!", zap.String("ok", "success"))
	InitMySql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
