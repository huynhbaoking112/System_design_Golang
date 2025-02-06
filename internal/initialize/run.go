package initialize

func Run() {
	InitLoadConfig()

	InitLogger()

	InitMySql()

	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
