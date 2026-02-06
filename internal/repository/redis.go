package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr        string
	Password    string
	User        string
	DB          int
	MaxRetries  int
	DialTimeout time.Duration
	RWTimeout   time.Duration
}

func NewRedisClient(ctx context.Context, config RedisConfig) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB,
		Username:     config.User,
		MaxRetries:   config.MaxRetries,
		DialTimeout:  config.DialTimeout,
		ReadTimeout:  config.RWTimeout,
		WriteTimeout: config.RWTimeout,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return db, nil
}
