package auth

import (
	serviceGrpc "github.com/kumarvikramshahi/grpc-server/pkg/auth/internal/grpc"
	"github.com/kumarvikramshahi/grpc-server/pkg/auth/internal/services"
	"google.golang.org/grpc"
)

func NewGrpcAuthServer(grpcServer *grpc.Server) {
	serviceGrpc.RegisterLogInServer(grpcServer, &services.LoginService{})
	serviceGrpc.RegisterSignUpServer(grpcServer, &services.SignUpService{})
}
