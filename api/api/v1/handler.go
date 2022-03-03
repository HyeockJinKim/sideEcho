package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"sideEcho/dto"
	"sideEcho/exchange"
)

//go:generate mockgen -package $GOPACKAGE -destination $PWD/api/api/v1/mock_$GOFILE sideEcho/api/api/v1 Handler

type Handler interface {
	buy(c *customContext) error
	sell(c *customContext) error
}

// Handler는 서버의 동작과 관련된 값들을 관리
type handler struct {
	controller exchange.Controller
}

func NewHandler(controller exchange.Controller) Handler {
	return &handler{
		controller: controller,
	}
}

// @Summary      Buy
// @Description  Buy value
// @Tags         Exchange
// @Accept       json
// @Produce      json
// @Param        req    body      dto.BuyRequest   true  "balance for buy"
// @Success      200    {object}  dto.BuyResponse
// @Failure      400    {object}  dto.ErrorResponse      "invalid request"
// @Failure      500    {object}  dto.ErrorResponse      "Internal error"
// @Router       /api/v1/buy      [post]
func (h *handler) buy(c *customContext) error {
	req := dto.BuyRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := h.controller.Buy(req.Value); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, dto.BuyResponse{Value: req.Value})
}

// @Summary      Sell
// @Description  Sell value
// @Tags         Exchange
// @Accept       json
// @Produce      json
// @Param        req    body      dto.SellRequest   true  "balance for sell"
// @Success      200    {object}  dto.SellResponse
// @Failure      400    {object}  dto.ErrorResponse       "invalid request"
// @Failure      500    {object}  dto.ErrorResponse       "Internal error"
// @Router       /api/v1/sell     [post]
func (h *handler) sell(c *customContext) error {
	req := dto.SellRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := h.controller.Sell(req.Value); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, dto.SellResponse{Value: req.Value})
}
