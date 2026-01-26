package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	go func() {
		if err := srv.Start("8080", handlers); err != nil {
			log.Fatalf("server startup failed: %s", err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %s", err.Error())
	}
}
