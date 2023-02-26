package config

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitCFG() error {
	viper.AddConfigPath("../../config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
