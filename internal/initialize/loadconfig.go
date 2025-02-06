package initialize

import (
	"fmt"

	"github.com/huynhbaoking112/System_design_Golang/global"
	"github.com/spf13/viper"
)

// type Config struct {
// 	Server struct {
// 		Port int `mapstructure:"port"`
// 	} `mapstructure:"server"`
// 	Databases []struct {
// 		User     string `mapstructure:"user"`
// 		Password string `mapstructure:"password"`
// 		Host     string `mapstructure:"host"`
// 	} `mapstructure:"databases"`
// }

func InitLoadConfig() {
	viper := viper.New()

	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	//read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	// read server configuration
	// fmt.Println("Server Port::", viper.GetInt("server.port"))
	// fmt.Println("Server Port::", viper.GetString("security.jwt.key"))
	// dir, _ := os.Getwd()
	// fmt.Println("Current working directory:", dir)

	//configure structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("Unable to decode configuration %v", err))
	}

}
