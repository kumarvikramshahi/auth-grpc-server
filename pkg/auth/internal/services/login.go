package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kumarvikramshahi/auth-grpc-server/configs"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/adaptor"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/auth/internal/grpc"
	"github.com/kumarvikramshahi/auth-grpc-server/pkg/domain"
	"github.com/redis/go-redis/v9"
)

type LoginService struct {
	grpc.UnimplementedLogInServer
	redisAdaptor adaptor.RedisAdaptor
}

func (service *LoginService) LogInUser(
	ctx context.Context, request *grpc.UserRequest,
) (*grpc.LoginResponse, error) {
	// check if user exists or not
	user, err := service.redisAdaptor.GetUser(ctx, request.Email)
	if err != nil && err != redis.Nil {
		log.Println("[services/LogInUser] error in getting user", err)
		customErr := status.Error(codes.Internal, err.Error())
		return nil, customErr
	}
	if err == redis.Nil {
		errResponse := status.Error(codes.Unauthenticated, domain.INVALID_EMAIL_PASS)
		return nil, errResponse
	}

	// if wrong pass
	inputHashedPass, err := HashPassword(request.Password)
	if err != nil {
		log.Println("[services/LogInUser] error in hashing pass", err)
		customErr := status.Error(codes.Internal, err.Error()+"- error in hashing pass")
		return nil, customErr
	}
	if inputHashedPass != user.Password {
		errResponse := status.Error(codes.Unauthenticated, domain.INVALID_EMAIL_PASS)
		return nil, errResponse
	}

	// if user exists with correct pass
	loginTime := time.Now().Add(time.Hour * 24).Unix()
	jwtToken, err := CreateToken(user.Email, loginTime)
	if err != nil {
		log.Println(err)
		customErr := status.Error(codes.Internal, err.Error()+"- error in creating JWT")
		return nil, customErr
	}

	return &grpc.LoginResponse{
		Response: &grpc.LoginResponse_Data{
			Data: &grpc.LoginSuccessResponse{
				Token:           jwtToken,
				ExpiryTimestamp: loginTime,
			},
		},
	}, nil
}

func HashPassword(message string) (string, error) {
	hasher := sha256.New()

	_, err := hasher.Write([]byte(message))
	if err != nil {
		return "", err
	}

	// apply salting
	hashedBytes := hasher.Sum(nil)

	// Converting byte slice to str
	hashedPassword := hex.EncodeToString(hashedBytes)

	return hashedPassword, nil
}

func CreateToken(email string, expiry int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":  email,
			"expiry": expiry,
		})

	tokenString, err := token.SignedString(configs.ServiceConfigs.AuthSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return configs.ServiceConfigs.AuthSecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
