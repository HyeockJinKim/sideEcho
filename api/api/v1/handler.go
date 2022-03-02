package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"sideEcho/dto"
	"sideEcho/exchange"
)

type Handler interface {
	buy(c *customContext) error
	sell(c *customContext) error
}

// Handler는 서버의 동작과 관련된 값들을 관리
type handler struct {
	manager exchange.Manager
}

func NewHandler(manager exchange.Manager) Handler {
	return &handler{
		manager: manager,
	}
}

func (h *handler) buy(c *customContext) error {
	req := dto.BuyRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := h.manager.Buy(req.Value); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, dto.BuyResponse{Value: req.Value})
}

func (h *handler) sell(c *customContext) error {
	req := dto.SellRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := h.manager.Sell(req.Value); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, dto.SellResponse{Value: req.Value})
}
