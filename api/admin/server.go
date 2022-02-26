package admin

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	prom := prometheus.NewPrometheus("echo", nil)
	e.Use(prom.HandlerFunc)
	prom.SetMetricsPath(e)
}
