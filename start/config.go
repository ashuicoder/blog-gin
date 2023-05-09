package start

import (
	"server/global"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yaml")
	//尝试进行配置读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(err)
	}

}
