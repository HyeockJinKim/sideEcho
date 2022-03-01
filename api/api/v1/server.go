package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"sideEcho/exchange"
	"sideEcho/stats"
)

type customHandler = func(ctx *customContext) error
type customMiddleware = func(handler customHandler) customHandler

/// 요청에서 사용되는 값들을 필드로 가지는 custom context
type customContext struct {
	echo.Context
	stats stats.Stats
}

func Route(e *echo.Group) {
	serverStats := stats.NewStats()
	exchangeManager := exchange.NewManager()
	h := NewHandler(exchangeManager)
	e.Use(wrapContextMiddleware(serverStats))
	e.POST("/buy", customWrapper(h.buy, requestStatMiddleware()))
	e.POST("/sell", customWrapper(h.sell, requestStatMiddleware()))
}

func wrapContextMiddleware(stats stats.Stats) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &customContext{
				Context: c,
				stats:   stats,
			}
			return next(cc)
		}
	}
}

func customWrapper(h customHandler, m ...customMiddleware) echo.HandlerFunc {
	return func(e echo.Context) error {
		if ctx, ok := e.(*customContext); ok {
			handler := h
			for _, middleware := range m {
				handler = middleware(handler)
			}
			return handler(ctx)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to wrap customContext")
	}
}

func requestStatMiddleware() customMiddleware {
	return func(next customHandler) customHandler {
		return func(c *customContext) error {
			c.stats.IncreaseRequestCount()
			err := next(c)
			if err != nil {
				c.stats.IncreaseFailureRequestCount()
				return err
			}
			c.stats.IncreaseSuccessRequestCount()
			return nil
		}
	}
}
