proto_auth:
	protoc \
	--go_out=pkg/auth/internal/grpc \
	--go_opt=paths=source_relative \
	--go-grpc_out=pkg/auth/internal/grpc \
	--go-grpc_opt=paths=source_relative \
	--proto_path=pkg/auth/internal/proto \
	pkg/auth/internal/proto/login.proto

	protoc \
	--go_out=pkg/auth/internal/grpc \
	--go_opt=paths=source_relative \
	--go-grpc_out=pkg/auth/internal/grpc \
	--go-grpc_opt=paths=source_relative \
	--proto_path=pkg/auth/internal/proto \
	pkg/auth/internal/proto/signup.proto