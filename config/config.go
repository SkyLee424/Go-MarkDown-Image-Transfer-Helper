package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig(path string) {
	// 使用 viper 管理配置文件
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("读取配置文件失败")
	}
}
