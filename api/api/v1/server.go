package v1

import (
	"net/http"

	"sideEcho/exchange"

	"sideEcho/stats"

	"github.com/labstack/echo/v4"
)

type customContext struct {
	echo.Context
	manager exchange.Manager
	stats   stats.Stats
}

func Route(e *echo.Group) {
	serverStats := stats.NewStats()
	exchangeManager := exchange.NewManager()
	e.Use(wrapContextMiddleware(exchangeManager, serverStats))
	e.GET("/buy", customContextWrapper(buy))
	e.GET("/sell", customContextWrapper(sell))
}

func wrapContextMiddleware(manager exchange.Manager, stats stats.Stats) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &customContext{
				Context: c,
				manager: manager,
				stats:   stats,
			}
			return next(cc)
		}
	}
}

func customContextWrapper(h func(c *customContext) error) echo.HandlerFunc {
	return func(e echo.Context) error {
		if ctx, ok := e.(*customContext); ok {
			ctx.stats.IncreaseRequestCount()
			return h(ctx)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to wrap customContext")
	}
}
