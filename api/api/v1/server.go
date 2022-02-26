package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"sideEcho/exchange"
	"sideEcho/stats"
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
	e.GET("/buy", customContextWrapper(requestStatMiddleware(buy)))
	e.GET("/sell", customContextWrapper(requestStatMiddleware(sell)))
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
			return h(ctx)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to wrap customContext")
	}
}

func requestStatMiddleware(h func(c *customContext) error) func(c *customContext) error {
	return func(c *customContext) error {
		c.stats.IncreaseRequestCount()
		err := h(c)
		if err != nil {
			c.stats.IncreaseFailureRequestCount()
			return err
		}
		c.stats.IncreaseSuccessRequestCount()
		return nil
	}
}
