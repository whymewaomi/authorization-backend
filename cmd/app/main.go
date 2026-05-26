package main

import (
	core_config "auth/internal/core/config"
	core_middleware "auth/internal/core/middleware"
	core_postgresql "auth/internal/core/repository/postgresql"
	core_redis "auth/internal/core/repository/redis"
	core_server "auth/internal/core/server"
	auth_repository "auth/internal/features/Auth/repostiory"
	auth_service "auth/internal/features/Auth/service"
	auth_transport_http "auth/internal/features/Auth/transport/http"
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

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

	app := core_server.NewApp(cfg)
	r := app.GetEngine()
	

	r.Use(gin.Recovery())
	r.Use(core_middleware.CORS())
	r.Use(gin.Logger())

  pool, err := core_postgresql.NewPool(ctx, fmt.Sprintf(
		"postgres://%s:%s@postgres:5432/auth?sslmode=disable", 
		cfg.PostgresUser,
		cfg.PostgresPassword,
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