package v3

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
)

type ListV3RolesByQuery struct {
	services *core.Services
	urls     map[string]string
}

func NewListV3RolesByQuery(services *core.Services, urls map[string]string) *ListV3RolesByQuery {
	return &ListV3RolesByQuery{services: services, urls: urls}
}

func (route *ListV3RolesByQuery) Register(e *echo.Echo) {
	e.GET("/v3/roles", route.Handle)
}

func (route *ListV3RolesByQuery) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusInternalServerError,
		"not implemented",
	)
}
