package adaptor

import (
	"context"
	"fmt"

	"github.com/kumarvikramshahi/auth-grpc-server/core"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/model"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/domain"
)

type RedisAdaptor struct {
}

func (redisAdaptor *RedisAdaptor) CreateUser(ctx context.Context, user model.User) error {
	// Validate the user struct
	if err := domain.Validator.Struct(user); err != nil {
		return fmt.Errorf("validation failed: %v", err)
	}

	// Store user in Redis
	_, err := core.RedisClient.HSet(ctx, "user:"+user.Email, map[string]interface{}{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	}).Result()
	return err
}

func (redisAdaptor *RedisAdaptor) GetUser(ctx context.Context, email string) (model.User, error) {
	var user model.User
	result, err := core.RedisClient.HGetAll(ctx, "user:"+email).Result()
	if err != nil {
		return user, err
	}

	if len(result) == 0 {
		return user, fmt.Errorf("user not found")
	}

	user.Name = result["name"]
	user.Email = result["email"]
	user.Password = result["password"]

	return user, nil
}
