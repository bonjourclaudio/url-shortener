package main

import (
	"github.com/claudioontheweb/url-shortener/config"
	"github.com/spf13/viper"
)


func main() {

	config.GetConfig()

	db_username := viper.GetString("DB_USERNAME")
	db_password := viper.GetString("DB_PASSWORD")
	db_name := viper.GetString("DB_NAME")
	server_port := viper.GetString("SERVER_PORT")
	mysql_port := viper.GetString("MYSQL_PORT")

	a := App{}
	a.Initialize(db_username, db_password, db_name, mysql_port)

	defer a.DB.Close()

	a.Run(server_port)

}