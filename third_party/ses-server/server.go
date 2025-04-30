package ses_server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	Port       string
	Host       string
	WebhookUrl string
	Database   PersistenceStorage
}

func NewRouter(db PersistenceStorage) *gin.Engine {

	router := gin.Default()

	router.POST("/v1/code/exchange", v1CodeExchange(db))
	router.POST("/v1/verify", v1Verify(db))
	router.GET("/v1/keys", v1GetKeys(db))

	return router
}

func Start(cfg ServerConfig) {
	if cfg.Database == nil {
		cfg.Database = newSimpleDatabase()
	}

	cfg.Database.Init()

	router := NewRouter(cfg.Database)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler: router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no params) by default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	cfg.Database.Shutdown()

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
