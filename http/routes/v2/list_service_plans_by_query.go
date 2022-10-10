package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
)

type ListServicePlansByQuery struct {
	services *core.Services
	urls     map[string]string
}

func NewListServicePlansByQuery(services *core.Services, urls map[string]string) *ListServicePlansByQuery {
	return &ListServicePlansByQuery{services: services, urls: urls}
}

func (route *ListServicePlansByQuery) Register(e *echo.Echo) {
	e.GET("/v2/service_plans", route.Handle)
}

func (route *ListServicePlansByQuery) Handle(c echo.Context) error {
	route.services.Logger(c)

	return c.JSON(
		http.StatusInternalServerError,
		"not implemented",
	)
}
