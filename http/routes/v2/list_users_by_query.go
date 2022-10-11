package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fcapi/core"
	"github.com/starkandwayne/fcapi/http/routes/registry"
)

type ListUsersByQuery struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewListUsersByQuery(services *core.Services, urls map[string]string) *ListUsersByQuery {
	return &ListUsersByQuery{
		path:     "/v2/users",
		verb:     registry.GET,
		ready:    false,
		services: services,
		urls:     urls,
	}
}

func (route *ListUsersByQuery) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *ListUsersByQuery) Path() string {
	return route.path
}

func (route *ListUsersByQuery) Ready() bool {
	return route.ready
}

func (route *ListUsersByQuery) Verb() registry.Verb {
	return route.verb
}

func (route *ListUsersByQuery) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusInternalServerError,
		"not implemented",
	)
}
