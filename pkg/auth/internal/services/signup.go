package services

import (
	"context"

	"github.com/kumarvikramshahi/grpc-server/pkg/auth/internal/adaptor"
	"github.com/kumarvikramshahi/grpc-server/pkg/auth/internal/grpc"
	"github.com/kumarvikramshahi/grpc-server/pkg/auth/internal/model"
	"github.com/kumarvikramshahi/grpc-server/pkg/domain"
	"github.com/redis/go-redis/v9"
)

type SignUpService struct {
	grpc.UnimplementedSignUpServer
	redisAdaptor adaptor.RedisAdaptor
}

func (service *SignUpService) ErrorResponse(message string) *grpc.SignUpResponse {
	return &grpc.SignUpResponse{
		Response: &grpc.SignUpResponse_Error{
			Error: &grpc.SignUpErrorResponse{
				Message: message,
			},
		},
	}
}

func (service *SignUpService) SignUpUser(
	ctx context.Context, request *grpc.RegisterUserRequest,
) (*grpc.SignUpResponse, error) {
	// check if user exists
	_, err := service.redisAdaptor.GetUser(ctx, request.Email)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if err != redis.Nil {
		errResponse := service.ErrorResponse(domain.USER_ALREADY_EXIST)
		return errResponse, nil
	}

	// if user doesn't exist
	var user model.User
	user.Email = request.Email
	user.Name = request.Name
	user.Password, err = HashPassword(request.Password)
	if err != nil {
		errResponse := service.ErrorResponse(domain.SERVER_ERROR)
		return errResponse, err
	}
	service.redisAdaptor.CreateUser(ctx, user)
	return &grpc.SignUpResponse{
		Response: &grpc.SignUpResponse_Data{
			Data: &grpc.SignUpSuccessResponse{
				Message: "User created",
			},
		},
	}, nil
}
