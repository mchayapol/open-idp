package main

import (
	"log"

	"rehatcher/openidp/server/config"
	"rehatcher/openidp/server/server"

	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	log.Printf("Config initialized")

	app := server.NewApp()

	log.Printf("Starting server on port %s...", viper.GetString("port"))

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
