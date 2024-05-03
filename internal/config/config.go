package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitViperEnv() {
	viper.AutomaticEnv()
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.Debug()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read .env %s", err)
	}
}
