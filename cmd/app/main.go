package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github/trad3r/go_temp.git/internal/handler"
	"github/trad3r/go_temp.git/internal/server"
	"github/trad3r/go_temp.git/internal/ws"
	"log"
	"os"
)

const (
	logFile = "./logs/log"
)

func init() {
	if err := getConfig(); err != nil {
		log.Fatal(err)
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}

}
func main() {
	logfile, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logrus.Fatal(err)
	}
	defer logfile.Close()

	hub := ws.NewHub(logfile)
	go hub.Run()

	h := handler.NewHandler(hub)
	s := new(server.Server)

	if err := s.Run(os.Getenv("PORT"), h.InitRoutes()); err != nil {
		hub.Logger.Fatal(err)
	}
}

func getConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
