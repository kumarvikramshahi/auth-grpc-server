package port

import (
	"context"

	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/grpc"
)

type ILoginService interface {
	LogInUser(
		ctx context.Context, request *grpc.UserRequest,
	) (*internal.LoginSuccess, error)
}

type ISignUpService interface {
	SignUpUser(
		ctx context.Context, request *grpc.RegisterUserRequest,
	) (*internal.SignUpSuccess, error)
}
