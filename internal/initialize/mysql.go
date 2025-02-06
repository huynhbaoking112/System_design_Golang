package initialize

import (
	"fmt"
	"time"

	"github.com/huynhbaoking112/System_design_Golang/global"
	"github.com/huynhbaoking112/System_design_Golang/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrorPanic(err, "InitMySql initialization error")
	global.Logger.Info("Initializing MySQL Successfully")
	global.Mdb = db

	// set Pool
	SetPool()

	migrateTables()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()

	if err != nil {
		global.Logger.Error("MySQL error setPool", zap.Error(err))
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleCons))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		global.Logger.Error("MySQL Migrating tables error", zap.Error(err))
	}
}
