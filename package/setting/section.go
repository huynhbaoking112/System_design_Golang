package setting

type Config struct {
	Mysql         MySQLSetting  `mapstructure:"mysql"`
	Server        Server        `mapstructure:"server"`
	LoggerSetting LoggerSetting `mapstructure:"log"`
	RedisSetting  RedisSetting  `mapstructure:"redis"`
}

type Server struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
	Poolsize int    `mapstructure:"poolsize"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_size      int    `mapstructure:"max_size"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Compress      bool   `mapstructure:"compress"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleCons     int    `mapstructure:"maxIdleCons"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}
