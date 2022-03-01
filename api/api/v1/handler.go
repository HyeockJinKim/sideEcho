package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"sideEcho/dto"
)

func buy(c *customContext) error {
	req := dto.BuyRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := c.manager.Buy(req.Value); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, dto.BuyResponse{Value: req.Value})
}

func sell(c *customContext) error {
	req := dto.SellRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := c.manager.Sell(req.Value); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, dto.SellResponse{Value: req.Value})
}
