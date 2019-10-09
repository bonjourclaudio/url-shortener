package main

import (
	"github.com/claudioontheweb/url-shortener/config"
	"github.com/spf13/viper"
)


func main() {

	config.GetConfig()

	db_username := viper.GetString("DB_USERNAME")
	db_password := viper.GetString("DB_USERNAME")
	db_name := viper.GetString("DB_NAME")
	port := viper.GetString("PORT")

	a := App{}
	a.Initialize(db_username, db_password, db_name)

	defer a.DB.Close()

	a.Run(port)

}