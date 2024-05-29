package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Cfg = new(AppConfig)

type AppConfig struct {
	Name          string `mapstructure:"name"`
	Mode          string `mapstructure:"mode"`
	Port          int    `mapstructure:"port"`
	Version       string `mapstructure:"version"`
	StartTime     string `mapstructure:"start_time"`
	MachineId     uint16 `mapstructure:"machine_id"`
	JwtKey        string `mapstructure:"jwt_key"`
	*MysqlConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
	*EtcdConfig   `mapstructure:"etcd"`
	*LogConfig    `mapstructure:"log"`
	*ServerConfig `mapstructure:"server"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	SSLMode      string `mapstructure:"sslmode"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	TimeZone     string `mapstructure:"time_zone"`
}

type RedisConfig struct {
	Host          string `mapstructure:"host"`
	Password      string `mapstructure:"password"`
	Port          int    `mapstructure:"port"`
	DB            int    `mapstructure:"db"`
	PoolSize      int    `mapstructure:"pool_size"`
	MinIdleConns  int    `mapstructure:"min_idle_conns"`
	ConnectType   string `mapstructure:"connect_type"`
	SSHRemoteHost string `mapstructure:"ssh_remote_host"`
	SSHFile       string `mapstructure:"ssh_file"`
}

type EtcdConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	LocalTime  bool   `mapstructure:"local_time"`
}

type ServerConfig struct {
	UserSrvAddress string `mapstructure:"user_srv_address"`
	PdfSrvAddress  string `mapstructure:"pdf_srv_address"`
}

func init() {
	var err error
	// read config
	viper.SetConfigName("gpdf")     // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err = viper.ReadInConfig()      // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// live watching and re-reading of config files
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		viper.Unmarshal(&Cfg)
	})

	// discover and read conf
	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err = viper.Unmarshal(&Cfg); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return
}
