package api

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4/middleware"

	"sideEcho/api/admin"

	"sideEcho/api/api"

	"github.com/labstack/echo/v4"

	"sideEcho/config"
)

type CleanupFunc func(ctx context.Context)

func externalAPI(cfg *config.Config) CleanupFunc {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Duration(cfg.Config.Api.TimeoutSec) * time.Second,
	}))

	apiRouter := e.Group("/api")
	api.Route(apiRouter)

	// Start server
	go func() {
		addr := fmt.Sprintf(":%d", cfg.Config.Api.Port)
		if err := e.Start(addr); err != nil {
			e.Logger.Fatal(err)
		}
	}()
	return func(ctx context.Context) {
		// graceful shutdown
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
		e.Logger.Warn("server down")
	}
}

func internalAPI(cfg *config.Config) CleanupFunc {
	e := echo.New()
	e.HideBanner = true
	admin.Route(e)
	go func() {
		addr := fmt.Sprintf(":%d", cfg.Config.Api.InternalPort)
		if err := e.Start(addr); err != nil {
			e.Logger.Fatal(err)
		}
	}()
	return func(ctx context.Context) {
		// graceful shutdown
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}
}

func NewAPI(cfg *config.Config) CleanupFunc {
	cleanupInternal := internalAPI(cfg)
	cleanupExternal := externalAPI(cfg)
	return func(ctx context.Context) {
		cleanupInternal(ctx)
		cleanupExternal(ctx)
	}
}
