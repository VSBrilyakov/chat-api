package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	chatApp "github.com/VSBrilyakov/chat-api"
	"github.com/VSBrilyakov/chat-api/internal/handler"
	"github.com/VSBrilyakov/chat-api/internal/repository"
	"github.com/VSBrilyakov/chat-api/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	pgdb, err := repository.NewPostgresDB(repository.Config{
		Host:     "postgres",
		Port:     "5432",
		Username: "admin",
		Password: os.Getenv("PG_PASSWORD"),
		DBName:   "chatdb",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	d, _ := pgdb.DB()
	defer d.Close()

	redisCfg := repository.RedisConfig{
		Addr:        "redis:6379",
		Password:    os.Getenv("REDIS_PASSWORD"),
		User:        "admin",
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 10 * time.Second,
		RWTimeout:   5 * time.Second,
	}

	rdb, err := repository.NewRedisClient(context.Background(), redisCfg)
	if err != nil {
		log.Fatalf("failed to initialize redis: %s", err.Error())
	}
	defer rdb.Close()

	repos := repository.NewRepository(pgdb, rdb)
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
