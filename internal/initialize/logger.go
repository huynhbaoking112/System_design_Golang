package initialize

import (
	"github.com/huynhbaoking112/System_design_Golang/global"
	"github.com/huynhbaoking112/System_design_Golang/package/logger"
	"go.uber.org/zap"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LoggerSetting)
	global.Logger.Info("Config Log Ok!!", zap.String("ok", "success"))
}

// 0932455836
