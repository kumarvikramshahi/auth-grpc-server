package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

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

func (service *LoginService) ErrorResponse(message string) *grpc.LoginResponse {
	return &grpc.LoginResponse{
		Response: &grpc.LoginResponse_Error{
			Error: &grpc.LoginErrorResponse{
				Message: message,
			},
		},
	}
}

func (service *LoginService) LogInUser(
	ctx context.Context, request *grpc.UserRequest,
) (*grpc.LoginResponse, error) {
	// check if user exists or not
	user, err := service.redisAdaptor.GetUser(ctx, request.Email)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if err == redis.Nil {
		errResponse := service.ErrorResponse(domain.INVALID_EMAIL_PASS)
		return errResponse, nil
	}

	// if wrong pass
	inputHashedPass, err := HashPassword(request.Password)
	if err != nil {
		errResponse := service.ErrorResponse(domain.SERVER_ERROR)
		return errResponse, err
	}
	if inputHashedPass != user.Password {
		errResponse := service.ErrorResponse(domain.INVALID_EMAIL_PASS)
		return errResponse, nil
	}

	// if user exists with correct pass
	loginTime := time.Now().Add(time.Hour * 24).Unix()
	jwtToken, err := CreateToken(user.Email, loginTime)
	if err != nil {
		log.Fatal(err)
		return nil, err
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
