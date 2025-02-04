package setting

type Config struct {
	Mysql  MySQLSetting `mapstructure:"mysql"`
	Server Server       `mapstructure:"server"`
}

type Server struct {
	Port string `mapstructure:"port"`
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
