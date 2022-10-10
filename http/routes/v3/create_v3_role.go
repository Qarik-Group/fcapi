package v3

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
)

type CreateV3Role struct {
	services *core.Services
	urls     map[string]string
}

func NewCreateV3Role(services *core.Services, urls map[string]string) *CreateV3Role {
	return &CreateV3Role{services: services, urls: urls}
}

func (route *CreateV3Role) Register(e *echo.Echo) {
	e.POST("/v3/roles", route.Handle)
}

func (route *CreateV3Role) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusCreated,
		nil,
	)
}
