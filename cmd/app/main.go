package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	_ "github.com/whymewaomi/authorization-backend/docs"
	core_config "github.com/whymewaomi/authorization-backend/internal/core/config"
	core_logger "github.com/whymewaomi/authorization-backend/internal/core/logger"
	core_middleware "github.com/whymewaomi/authorization-backend/internal/core/middleware"
	core_postgresql "github.com/whymewaomi/authorization-backend/internal/core/repository/postgresql"
	core_redis "github.com/whymewaomi/authorization-backend/internal/core/repository/redis"
	core_server "github.com/whymewaomi/authorization-backend/internal/core/server"
	auth_repository "github.com/whymewaomi/authorization-backend/internal/features/Auth/repostiory"
	auth_service "github.com/whymewaomi/authorization-backend/internal/features/Auth/service"
	auth_transport_http "github.com/whymewaomi/authorization-backend/internal/features/Auth/transport/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Authorization API
// @version 1.0
// @description Authentication backend with JWT
// @host localhost:5050
// @BasePath /
func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	cfg := core_config.Load()

	if err := cfg.Validation(); err != nil {
		log.Fatalf("failed load config: %v", err)
	}

	logger, file, err := core_logger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	app := core_server.NewApp(cfg)
	r := app.GetEngine()

	r.Use(gin.Recovery())
	r.Use(core_middleware.RequestID())
	r.Use(core_middleware.CORS())
	r.Use(core_middleware.Logger(logger))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pool, err := core_postgresql.NewPool(ctx, fmt.Sprintf(
		"postgres://%s:%s@postgres:5432/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
	))
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	rdb := core_redis.NewClient(cfg)
	defer rdb.Close()

	redisStorage := core_redis.NewRedisStorage(rdb)

	authRepository := auth_repository.NewAuthRepository(pool)
	authService := auth_service.NewAuthService(authRepository, redisStorage)
	authTransport := auth_transport_http.NewAuth(authService, r)

	authTransport.RegisterRouter()

	if err := app.Run(ctx); err != nil {
		log.Fatalf("failed run HTTP server: %v", err)
	}
}
