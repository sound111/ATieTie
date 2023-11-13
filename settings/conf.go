package settings

// 注意tag的引号！！！！

type Config struct {
	*AppConfig   `mapstructure:"app"`
	*SnowConfig  `mapstructure:"snow"`
	*LogConfig   `mapstructure:"log"`
	*MYSQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type SnowConfig struct {
	StartTime string `mapstructure:"start_time"`
	MachineID uint16 `mapstructure:"machine_id"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
}

type MYSQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Pwd          string `mapstructure:"pwd"`
	DB           string `mapstructure:"db"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Pwd  string `mapstructure:"pwd"`
	DB   int    `mapstructure:"db"`
	//PoolSize int `mapstructure:"pool_size"`
}
