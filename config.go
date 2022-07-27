package main

import (
	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}

func getKey() string {
	return viper.GetString("SNOB_KEY")
}

func setKey(key string) {
	viper.Set("SNOB_KEY", key)
	viper.WriteConfig()
}
