package auth

import (
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/adaptor"
	serviceGrpc "github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/grpc"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/port"
	"google.golang.org/grpc"
)

func NewGrpcAuthServer(grpcServer *grpc.Server) {
	redisAdaptor := adaptor.NewRedisAdaptor()

	clientPort := port.NewAuthClientPort(redisAdaptor)

	serviceGrpc.RegisterLogInServer(grpcServer, clientPort)
	serviceGrpc.RegisterSignUpServer(grpcServer, clientPort)
}
