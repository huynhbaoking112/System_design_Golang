package initialize

func Run() {
	InitLoadConfig()
	// fmt.Println("Loading configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	InitMySql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
