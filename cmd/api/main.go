package main

import (
	"log"
	"log/slog"

	"github.com/cukhoaimon/khoainats/internal/auth"
	"github.com/cukhoaimon/khoainats/internal/logger"
	"github.com/cukhoaimon/khoainats/internal/resource"
	"github.com/cukhoaimon/khoainats/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Repository ########################################

	// Services ##########################################
	authService := services.NewAuthService(slog.LevelInfo)

	// Resources ##########################################
	apiResource := resource.NewDefaultAPI(
		resource.NewDefaultAPIConfig{
			AuthService: authService,
			Middlewares: []gin.HandlerFunc{
				auth.X5009AuthFilter(),
				gin.Recovery(),
				gin.LoggerWithFormatter(logger.CustomGinLogger),
			},
			LogLevel: slog.LevelInfo,
		},
	)

	log.Fatal(apiResource.Run("localhost:8080"))
}
