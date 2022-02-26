package api

import (
	"github.com/labstack/echo/v4"

	v1 "sideEcho/api/api/v1"
)

func Route(group *echo.Group) {
	v1Router := group.Group("/v1")
	v1.Route(v1Router)
}
