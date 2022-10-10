package v3

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
	"codeberg.org/ess/fcapi/http/routes/registry"
)

type CreateV3Role struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewCreateV3Role(services *core.Services, urls map[string]string) *CreateV3Role {
	return &CreateV3Role{
		path:     "/v3/roles",
		verb:     registry.POST,
		ready:    true,
		services: services,
		urls:     urls,
	}
}

func (route *CreateV3Role) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *CreateV3Role) Path() string {
	return route.path
}

func (route *CreateV3Role) Ready() bool {
	return route.ready
}

func (route *CreateV3Role) Verb() registry.Verb {
	return route.verb
}

func (route *CreateV3Role) Handle(c echo.Context) error {
	route.services.Logger(c)

	//TODO: While this is enough to make the cf cli work, we should probably
	// actually store the role.
	return c.JSON(
		http.StatusCreated,
		nil,
	)
}
