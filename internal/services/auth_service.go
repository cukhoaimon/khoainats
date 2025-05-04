package services

import (
	"log"
	"log/slog"

	"github.com/cukhoaimon/khoainats/api/generated/server"
	"github.com/cukhoaimon/khoainats/internal/logger"
	"github.com/cukhoaimon/khoainats/pkg/ses"
)

type AuthService interface {
	V1LoginStart(request openapi.V1LoginStartRequest) (openapi.V1LoginStartResponse, error)
	V1LoginExchange(request openapi.V1LoginExchangeRequest) (openapi.V1AccessToken, error)
}

type defaultAuthService struct {
	logger    *log.Logger
	sesClient ses.Client
}

func NewAuthService(logLevel slog.Level, sesClient ses.Client) AuthService {
	return defaultAuthService{
		logger:    logger.New(logLevel),
		sesClient: sesClient,
	}
}

// ------

func (s defaultAuthService) V1LoginStart(request openapi.V1LoginStartRequest) (openapi.V1LoginStartResponse, error) {
	// handle logic
	s.logger.Printf("V1LoginStart request: %+v", request)
	resp, err := s.sesClient.V1CodeExchange(request.Email)
	if err != nil {
		return openapi.V1LoginStartResponse{}, nil
	}

	return openapi.V1LoginStartResponse{}, nil
}

func (s defaultAuthService) V1LoginExchange(request openapi.V1LoginExchangeRequest) (openapi.V1AccessToken, error) {
	s.logger.Printf("V1LoginExchange request: %+v", request)
	return openapi.V1AccessToken{}, nil
}
