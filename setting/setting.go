package setting

import "gopkg.in/ini.v1"

// MySQLConfig 数据库配置信息
type MySQLConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

// RedisConfig Redis配置信息
type RedisConfig struct {
	RedisAddress  string `json:"address"`
	RedisPassword string `json:"password"`
	RedisDB       int    `json:"db"`
}

// AppConfig 应用配置信息
type AppConfig struct {
	Port         int  `json:"port"`
	Release      bool `json:"release"`
	*MySQLConfig `json:"database"`
	*RedisConfig `json:"redis"`
}

var (
	Conf = new(AppConfig)
)

// Init 加载配置文件
func Init(file string) error {
	cfg, err := ini.Load(file)
	if err != nil {
		return err
	}
	Conf.Port = cfg.Section("").Key("port").MustInt(8080)
	Conf.Release = cfg.Section("").Key("release").MustBool(false)

	Conf.MySQLConfig = &MySQLConfig{
		Host:     cfg.Section("database").Key("host").String(),
		Port:     cfg.Section("database").Key("port").MustInt(3306),
		User:     cfg.Section("database").Key("user").String(),
		Password: cfg.Section("database").Key("password").String(),
		DBName:   cfg.Section("database").Key("dbname").String(),
	}
	Conf.RedisConfig = &RedisConfig{
		RedisAddress:  cfg.Section("redis").Key("address").String(),
		RedisPassword: cfg.Section("redis").Key("password").String(),
		RedisDB:       cfg.Section("redis").Key("db").MustInt(0),
	}
	return err
}
