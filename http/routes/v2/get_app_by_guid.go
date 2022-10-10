package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
)

type GetAppByGuid struct {
	services *core.Services
	urls     map[string]string
}

func NewGetAppByGuid(services *core.Services, urls map[string]string) *GetAppByGuid {
	return &GetAppByGuid{services: services, urls: urls}
}

func (route *GetAppByGuid) Register(e *echo.Echo) {
	e.GET("/v2/apps/:guid", route.Handle)
}

func (route *GetAppByGuid) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusInternalServerError,
		"not implemented",
	)
}
