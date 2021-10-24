package router

import (
	"auto-booking/handler"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	AutoBookingHandler handler.AutoBookingHandler
}

func (api *API) SetupRouter() {
	api.Echo.GET("/", handler.Welcome)
}
