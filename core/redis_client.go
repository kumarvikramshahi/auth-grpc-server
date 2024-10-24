package core

import (
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
}
