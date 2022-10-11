package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fcapi/core"
	"github.com/starkandwayne/fcapi/http/routes/registry"
)

type GetAppByGuid struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewGetAppByGuid(services *core.Services, urls map[string]string) *GetAppByGuid {
	return &GetAppByGuid{
		path:     "/v2/apps/:guid",
		verb:     registry.GET,
		ready:    true,
		services: services,
		urls:     urls,
	}
}

func (route *GetAppByGuid) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *GetAppByGuid) Path() string {
	return route.path
}

func (route *GetAppByGuid) Verb() registry.Verb {
	return route.verb
}

func (route *GetAppByGuid) Ready() bool {
	return route.ready
}

func (route *GetAppByGuid) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusInternalServerError,
		"not implemented",
	)
}
