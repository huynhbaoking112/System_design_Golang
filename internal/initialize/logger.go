package initialize

import (
	"github.com/huynhbaoking112/System_design_Golang/global"
	"github.com/huynhbaoking112/System_design_Golang/package/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LoggerSetting)
}
