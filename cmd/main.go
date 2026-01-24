package main

import (
	"log"

	"github.com/VSBrilyakov/chat-api"
	"github.com/VSBrilyakov/chat-api/internal/handler"
	"github.com/VSBrilyakov/chat-api/internal/repository"
	"github.com/VSBrilyakov/chat-api/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "postgres",
		Port:     "5432",
		Username: "admin",
		Password: "pgpassword123",
		DBName:   "chatdb",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHTTPHandler(services)
	handlers.InitRoutes()

	srv := new(chatApp.Server)
	if err := srv.Start("8080", handlers); err != nil {
		log.Fatalf("server startup failed: %s", err.Error())
	}

}
