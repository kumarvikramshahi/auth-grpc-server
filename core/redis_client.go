package core

import (
	"context"
	"fmt"
	"log"

	"github.com/kumarvikramshahi/auth-grpc-server/configs"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func NewRedisSingletonClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     configs.ServiceConfigs.RedisConfigs.Uri,
		Password: configs.ServiceConfigs.RedisConfigs.Password,
		DB:       configs.ServiceConfigs.RedisConfigs.Database,
		Protocol: configs.ServiceConfigs.RedisConfigs.Protocol, // Connection protocol
	})

	ctx := context.Background()
	pong, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	} else {
		fmt.Println("Redis connection status:", pong)
	}
}
