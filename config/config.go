package config

import "github.com/spf13/viper"

func GetConfig() {

	viper.SetConfigFile(`./config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}