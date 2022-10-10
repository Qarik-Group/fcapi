package v3

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
)

type GetTaskByGuid struct {
	services *core.Services
	urls     map[string]string
}

func NewGetTaskByGuid(services *core.Services, urls map[string]string) *GetTaskByGuid {
	return &GetTaskByGuid{services: services, urls: urls}
}

func (route *GetTaskByGuid) Register(e *echo.Echo) {
	e.GET("/v3/tasks/:guid", route.Handle)
}

func (route *GetTaskByGuid) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusInternalServerError,
		"not implemented",
	)
}
