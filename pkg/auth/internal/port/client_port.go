package port

import (
	"context"

	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/adaptor"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/grpc"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthClientPort struct {
	grpc.UnimplementedLogInServer
	grpc.UnimplementedSignUpServer
	ILoginService  ILoginService
	ISignUpService ISignUpService
}

func NewAuthClientPort(redisAdaptor *adaptor.RedisAdaptor) *AuthClientPort {
	return &AuthClientPort{
		ILoginService:  services.NewLoginService(redisAdaptor),
		ISignUpService: services.NewSignUpService(redisAdaptor),
	}
}

func (service *AuthClientPort) LogInUser(
	ctx context.Context, request *grpc.UserRequest,
) (*grpc.LoginResponse, error) {

	// TODO => create sepearate methods for  req validation
	// temporary validate methods
	if request.Email == "" || request.Password == "" {
		customErr := status.Error(
			codes.InvalidArgument,
			"either email or password is empty/incorrect_field_name",
		)
		return nil, customErr
	}

	loginCheck, err := service.ILoginService.LogInUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return &grpc.LoginResponse{
		Response: &grpc.LoginResponse_Data{
			Data: &grpc.LoginSuccessResponse{
				Token:           loginCheck.Token,
				ExpiryTimestamp: loginCheck.ExpiryTimestamp,
			},
		},
	}, nil
}

func (service *AuthClientPort) SignUpUser(
	ctx context.Context, request *grpc.RegisterUserRequest,
) (*grpc.SignUpResponse, error) {

	// request validation
	if request.Email == "" || request.Password == "" || request.Name == "" {
		customErr := status.Error(
			codes.InvalidArgument,
			"either email or password or name is empty/incorrect_field_name",
		)
		return nil, customErr
	}

	signupCheck, err := service.ISignUpService.SignUpUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return &grpc.SignUpResponse{
		Response: &grpc.SignUpResponse_Data{
			Data: &grpc.SignUpSuccessResponse{
				Message: signupCheck.Message,
			},
		},
	}, nil
}
