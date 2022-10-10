package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
)

type ListUsersByQuery struct {
	services *core.Services
	urls     map[string]string
}

func NewListUsersByQuery(services *core.Services, urls map[string]string) *ListUsersByQuery {
	return &ListUsersByQuery{services: services, urls: urls}
}

func (route *ListUsersByQuery) Register(e *echo.Echo) {
	e.GET("/v2/users", route.Handle)
}

func (route *ListUsersByQuery) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusInternalServerError,
		"not implemented",
	)
}
