package v3

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
)

type CreateTask struct {
	services *core.Services
	urls     map[string]string
}

func NewCreateTask(services *core.Services, urls map[string]string) *CreateTask {
	return &CreateTask{services: services, urls: urls}
}

func (route *CreateTask) Register(e *echo.Echo) {
	e.POST("/v3/apps/:guid/tasks", route.Handle)
}

func (route *CreateTask) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusInternalServerError,
		"not implemented",
	)
}
