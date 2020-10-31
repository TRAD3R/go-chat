package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github/trad3r/go_temp.git/internal/handler"
	"github/trad3r/go_temp.git/internal/repository"
	"github/trad3r/go_temp.git/internal/server"
	"github/trad3r/go_temp.git/internal/service"
	"log"
	"os"
)

func init() {
	if err := getConfig(); err != nil {
		log.Fatal(fmt.Sprintf("Error read config.yml: %v", err.Error()))
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(fmt.Sprintf("Error read .env: %v", err.Error()))
	}
}
func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatal(fmt.Sprintf("Connect to db error: %v", err.Error()))
	}

	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	h := handler.NewHandler(serv)
	s := new(server.Server)
	if err := s.Run(viper.GetString("port"), h.InitRoutes()); err != nil {
		log.Fatal(fmt.Sprintf("Run server error: %v", err.Error()))
	}
}

func getConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
