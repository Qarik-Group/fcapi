package v3

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fcapi/core"
	"github.com/starkandwayne/fcapi/http/routes/registry"
)

type CreateTask struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewCreateTask(services *core.Services, urls map[string]string) *CreateTask {
	return &CreateTask{
		path:     "/v3/apps/:guid/tasks",
		verb:     registry.POST,
		ready:    false,
		services: services,
		urls:     urls,
	}
}

func (route *CreateTask) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *CreateTask) Path() string {
	return route.path
}

func (route *CreateTask) Ready() bool {
	return route.ready
}

func (route *CreateTask) Verb() registry.Verb {
	return route.verb
}

func (route *CreateTask) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusInternalServerError,
		"not implemented",
	)
}
