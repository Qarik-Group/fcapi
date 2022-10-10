package fcapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
	"codeberg.org/ess/fcapi/http/routes/registry"
)

type GetRoutes struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewGetRoutes(services *core.Services, urls map[string]string) *GetRoutes {
	return &GetRoutes{
		path:     "/_fcapi/routes",
		verb:     registry.GET,
		ready:    true,
		services: services,
		urls:     urls,
	}
}

func (route *GetRoutes) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *GetRoutes) Path() string {
	return route.path
}

func (route *GetRoutes) Ready() bool {
	return route.ready
}

func (route *GetRoutes) Verb() registry.Verb {
	return route.verb
}

func (route *GetRoutes) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusOK,
		registry.Registered(),
	)
}
