package admin

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Route(e *echo.Echo) {
	prom := prometheus.NewPrometheus("echo", nil)
	e.Use(prom.HandlerFunc)
	prom.SetMetricsPath(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
