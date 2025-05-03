package services

import (
	"log"
	"log/slog"

	openapi "github.com/cukhoaimon/khoainats/api/generated/server"
	"github.com/cukhoaimon/khoainats/internal/logger"
)

type AuthService interface {
	V1LoginStart(request openapi.V1LoginStartRequest) (openapi.V1LoginStartResponse, error)
	V1LoginExchange(request openapi.V1LoginExchangeRequest) (openapi.V1AccessToken, error)
}

type defaultAuthService struct {
	logger *log.Logger
}

func NewAuthService(logLevel slog.Level) AuthService {
	return defaultAuthService{
		logger: logger.New(logLevel),
	}
}

// ------

func (s defaultAuthService) V1LoginStart(request openapi.V1LoginStartRequest) (openapi.V1LoginStartResponse, error) {
	// handle logic
	s.logger.Printf("V1LoginStart request: %+v", request)
	return openapi.V1LoginStartResponse{}, nil
}

func (s defaultAuthService) V1LoginExchange(request openapi.V1LoginExchangeRequest) (openapi.V1AccessToken, error) {
	s.logger.Printf("V1LoginExchange request: %+v", request)
	return openapi.V1AccessToken{}, nil
}
