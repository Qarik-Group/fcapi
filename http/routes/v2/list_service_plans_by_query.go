package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fcapi/core"
	"github.com/starkandwayne/fcapi/http/routes/registry"
)

type ListServicePlansByQuery struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewListServicePlansByQuery(services *core.Services, urls map[string]string) *ListServicePlansByQuery {
	return &ListServicePlansByQuery{
		path:     "/v2/service_plans",
		verb:     registry.GET,
		ready:    false,
		services: services,
		urls:     urls,
	}
}

func (route *ListServicePlansByQuery) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *ListServicePlansByQuery) Path() string {
	return route.path
}

func (route *ListServicePlansByQuery) Ready() bool {
	return route.ready
}

func (route *ListServicePlansByQuery) Verb() registry.Verb {
	return route.verb
}

func (route *ListServicePlansByQuery) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusInternalServerError,
		"not implemented",
	)
}
