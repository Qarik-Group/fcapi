package v3

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fcapi/core"
	"github.com/starkandwayne/fcapi/http/routes/registry"
)

type GetTaskByGuid struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewGetTaskByGuid(services *core.Services, urls map[string]string) *GetTaskByGuid {
	return &GetTaskByGuid{
		path:     "/v3/tasks/:guid",
		verb:     registry.GET,
		ready:    false,
		services: services,
		urls:     urls,
	}
}

func (route *GetTaskByGuid) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *GetTaskByGuid) Path() string {
	return route.path
}

func (route *GetTaskByGuid) Ready() bool {
	return route.ready
}

func (route *GetTaskByGuid) Verb() registry.Verb {
	return route.verb
}

func (route *GetTaskByGuid) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusInternalServerError,
		"not implemented",
	)
}
