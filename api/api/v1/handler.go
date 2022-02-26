package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"sideEcho/dto"
)

func hello(c *customContext) error {
	return c.String(http.StatusOK, "hello")
}

func buy(c *customContext) error {
	req := new(dto.SellRequest)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := c.manager.Sell(req.Value); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "buy success")
}

func sell(c *customContext) error {
	req := new(dto.SellRequest)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := c.manager.Sell(req.Value); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "sell success")
}
