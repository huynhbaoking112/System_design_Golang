package global

import (
	"github.com/huynhbaoking112/System_design_Golang/package/logger"
	"github.com/huynhbaoking112/System_design_Golang/package/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
