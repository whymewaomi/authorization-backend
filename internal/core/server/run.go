package core_server

import (
	core_config "auth/internal/core/config"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	cfg *core_config.Config

	engine *gin.Engine
}

func NewApp(
	cfg *core_config.Config,
) *App {
	return &App{
		cfg: cfg,
		engine: gin.New(),
	}
}

func (a *App) GetEngine() *gin.Engine {
	return a.engine
}

func (a *App) Run(ctx context.Context) error {
	server := http.Server{
		Addr: fmt.Sprintf(":%s", a.cfg.Host),
		Handler: a.engine,
	}

	ch := make(chan error, 1)

	go func() {
    
		log.Printf("HTTP server started...")
		log.Printf("HTTP server started in addr :%s", a.cfg.Host)

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
			return 
		}

		ch <- nil
	}()

	select {
	case err := <- ch:
		if err != nil {
			return fmt.Errorf("failed HTTP started server: %w", err)
		}
	case <- ctx.Done():
		log.Printf("Shutdown HTTP server...")
		
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("failed shutdown HTTP server: %v", err)
		}

		return nil
	}

	return nil
}