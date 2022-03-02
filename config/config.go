package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*Path        `mapstructure:"path"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}
type Path struct {
	Logs string `mapstructure:"logs"`
	Data string `mapstructure:"data"`
	Temp string `mapstructure:"temp"`
}

func Init() (err error) {

	viper.SetConfigFile("./etc/config.yml")
	viper.AddConfigPath(".")   // 指定查找配置文件的路径（这里使用相对路径）
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}
	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
	// 监控配置文件修改
	//viper.WatchConfig()
	//viper.OnConfigChange(func(in fsnotify.Event) {
	//	fmt.Println("配置文件修改了...")
	//	if err := viper.Unmarshal(Conf); err != nil {
	//		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	//	}
	//})
	// 创建目录
	err = os.MkdirAll(Conf.Path.Temp+"/img/identicon", 777)
	if err != nil {
		fmt.Println(err)
	}
	err = os.MkdirAll(Conf.Path.Logs, 777)
	if err != nil {
		fmt.Println(err)
	}
	return
}
