package resource

import (
	"log"
	"log/slog"
	"net/http"

	openapi "github.com/cukhoaimon/khoainats/api/generated/server"
	"github.com/cukhoaimon/khoainats/internal/auth"
	"github.com/cukhoaimon/khoainats/internal/logger"
	"github.com/cukhoaimon/khoainats/internal/services"
	"github.com/gin-gonic/gin"
)

type handlers struct {
	authService services.AuthService
	logger      *log.Logger
}

func (h *handlers) V1LoginStart(c *gin.Context) {
	baseHandler[openapi.V1LoginStartRequest, openapi.V1LoginStartResponse](
		c,
		func(req openapi.V1LoginStartRequest) (openapi.V1LoginStartResponse, error) {
			return h.authService.V1LoginStart(req)
		},
	)
}

func (h *handlers) V1LoginExchange(c *gin.Context) {
	baseHandler[openapi.V1LoginExchangeRequest, openapi.V1AccessToken](
		c,
		func(req openapi.V1LoginExchangeRequest) (openapi.V1AccessToken, error) {
			return h.authService.V1LoginExchange(req)
		},
	)
}

func baseHandler[Req any, Res any](
	c *gin.Context,
	action func(req Req) (Res, error),
) {
	var req Req

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := action(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusAccepted, res)
}

// ---------------------------

type NewDefaultAPIConfig struct {
	AuthService services.AuthService
	Middlewares []gin.HandlerFunc
	LogLevel    slog.Level
}

func NewDefaultAPI(cfg NewDefaultAPIConfig) *gin.Engine {
	engine := gin.New()

	for _, middleware := range cfg.Middlewares {
		engine.Use(middleware)
	}

	engineWithRouter := openapi.NewRouterWithGinEngine(
		engine,
		auth.JwtRequestFilter,
		&handlers{
			logger:      logger.New(cfg.LogLevel),
			authService: cfg.AuthService,
		},
	)

	return engineWithRouter
}

// ---------------------------
