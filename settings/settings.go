package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	//viper使用的tag是mapstructure
)

var Conf *Config

func Init(f string) (err error) {
	viper.SetConfigFile(f)
	//f 配置文件路径

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		return
	}

	//热加载，监控配置文件的变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file change:", e.Name)

		err = viper.Unmarshal(&Conf)
		if err != nil {
			return
		}
	})

	return
}
