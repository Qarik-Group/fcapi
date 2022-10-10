package v3

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
	"codeberg.org/ess/fcapi/http/responses"
	"codeberg.org/ess/fcapi/http/routes/registry"
)

type ListV3OrganizationsByQuery struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewListV3OrganizationsByQuery(services *core.Services, urls map[string]string) *ListV3OrganizationsByQuery {
	return &ListV3OrganizationsByQuery{
		path:     "/v3/organizations",
		verb:     registry.GET,
		ready:    true,
		services: services,
		urls:     urls,
	}
}

func (route *ListV3OrganizationsByQuery) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *ListV3OrganizationsByQuery) Path() string {
	return route.path
}

func (route *ListV3OrganizationsByQuery) Ready() bool {
	return route.ready
}

func (route *ListV3OrganizationsByQuery) Verb() registry.Verb {
	return route.verb
}

func (route *ListV3OrganizationsByQuery) Handle(c echo.Context) error {
	route.services.Logger(c)

	orgs := route.services.Orgs.All()

	output := &responses.OrgCollection{
		Resources: orgs,
		Pagination: &responses.Pagination{
			TotalPages:   1,
			TotalResults: len(orgs),
			First:        nil,
			Last:         nil,
			Next:         nil,
			Previous:     nil,
		},
	}

	return c.JSON(
		http.StatusOK,
		output,
	)
}
