package main

import (
	"github.com/spf13/viper"
	"github/trad3r/go_temp.git/internal/handler"
	"github/trad3r/go_temp.git/internal/ws"
	"log"
)

func init() {
	if err := getConfig(); err != nil {
		log.Fatal(err)
	}
}
func main() {
	hub := ws.NewHub()
	go hub.Run()

	h := handler.NewHandler(hub)
	if err := h.InitRoutes().Run(":" + viper.GetString("port")); err != nil {
		log.Fatal(err)
	}
}

func getConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
