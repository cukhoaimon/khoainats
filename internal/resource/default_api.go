package resource

import (
	"fmt"
	"net/http"

	openapi "github.com/cukhoaimon/khoainats/api/generated/server"
	"github.com/cukhoaimon/khoainats/internal/auth"
	"github.com/cukhoaimon/khoainats/internal/services"
	"github.com/gin-gonic/gin"
)

type handlers struct {
	authService services.AuthService
}

func (h handlers) V1NoauthGet(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h handlers) V1LoginStart(c *gin.Context) {
	var req openapi.V1LoginStartRequest
	var err error

	if err = c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.ErrBodyNotAllowed)
	}

	response, err := h.authService.CreateLogin(req.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}

	c.JSON(http.StatusAccepted, openapi.EmptySuccess{Message: response})
}

// ---------------------------

type NewDefaultAPIConfig struct {
	AuthService services.AuthService
	Middlewares []gin.HandlerFunc
}

func NewDefaultAPI(cfg NewDefaultAPIConfig) *gin.Engine {
	engine := gin.New()

	for _, middleware := range cfg.Middlewares {
		engine.Use(middleware)
	}

	engineWithRouter := openapi.NewRouterWithGinEngine(
		engine,
		auth.JwtRequestFilter,
		handlers{
			authService: cfg.AuthService,
		},
	)

	return engineWithRouter
}

// ---------------------------
