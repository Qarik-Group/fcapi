package v3

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
	"codeberg.org/ess/fcapi/http/routes/registry"
)

type ListV3RolesByQuery struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewListV3RolesByQuery(services *core.Services, urls map[string]string) *ListV3RolesByQuery {
	return &ListV3RolesByQuery{
		path:     "/v3/roles",
		verb:     registry.GET,
		ready:    false,
		services: services,
		urls:     urls,
	}
}

func (route *ListV3RolesByQuery) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *ListV3RolesByQuery) Path() string {
	return route.path
}

func (route *ListV3RolesByQuery) Ready() bool {
	return route.ready
}

func (route *ListV3RolesByQuery) Verb() registry.Verb {
	return route.verb
}

func (route *ListV3RolesByQuery) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusInternalServerError,
		"not implemented",
	)
}
