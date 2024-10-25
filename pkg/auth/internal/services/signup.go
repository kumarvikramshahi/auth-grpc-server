package services

import (
	"context"
	"log"

	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/adaptor"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/grpc"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/model"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/domain"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SignUpService struct {
	redisAdaptor *adaptor.RedisAdaptor
}

func NewSignUpService(redisAdaptor *adaptor.RedisAdaptor) *SignUpService {
	return &SignUpService{
		redisAdaptor: redisAdaptor,
	}
}

func (service *SignUpService) SignUpUser(
	ctx context.Context, request *grpc.RegisterUserRequest,
) (*internal.SignUpSuccess, error) {
	// check if user exists
	_, err := service.redisAdaptor.GetUser(ctx, request.Email)
	if err != nil && err != redis.Nil {
		log.Println("[services/SignUpUser] error in fetching from redis - ", err)
		customErr := status.Error(codes.Internal, err.Error()+" - error in fetching from redis")
		return nil, customErr
	}
	if err != redis.Nil {
		errResponse := status.Error(codes.AlreadyExists, domain.USER_ALREADY_EXIST)
		return nil, errResponse
	}

	// if user doesn't exist
	var user model.User
	user.Email = request.Email
	user.Name = request.Name
	user.Password, err = HashPassword(request.Password)
	if err != nil {
		log.Println("[services/LogInUser] error in hashing pass", err)
		customErr := status.Error(codes.Internal, err.Error()+"- error hashing pass")
		return nil, customErr
	}
	err = service.redisAdaptor.CreateUser(ctx, user)
	if err != nil {
		log.Println("[services/LogInUser] error in creating user", err)
		customErr := status.Error(codes.Internal, err.Error()+"- error in creating user")
		return nil, customErr
	}
	return &internal.SignUpSuccess{
		Message: "User created",
	}, nil
}
