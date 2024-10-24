package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/kumarvikramshahi/grpc-server/configs"
	"github.com/kumarvikramshahi/grpc-server/core"
	"github.com/kumarvikramshahi/grpc-server/pkg/auth"
	"github.com/kumarvikramshahi/grpc-server/pkg/domain"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	// Load Configuration File
	settingsFileName := os.Args[1]
	configs.LoadServiceConfigurations(settingsFileName)

	// validator
	domain.ValidatorSingletonClient()

	// init Dbs
	core.NewRedisSingletonClient()

	listen, err := net.Listen("tcp", ":"+configs.ServiceConfigs.Port)
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	log.Println("server started")

	// server
	serverRegistrar := grpc.NewServer()

	// Register reflection service on gRPC server
	reflection.Register(serverRegistrar)

	// services
	auth.NewGrpcAuthServer(serverRegistrar)

	err = serverRegistrar.Serve(listen)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
