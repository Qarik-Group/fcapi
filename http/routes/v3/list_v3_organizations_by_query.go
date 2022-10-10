package v3

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
	"codeberg.org/ess/fcapi/http/responses"
)

type ListV3OrganizationsByQuery struct {
	services *core.Services
	urls     map[string]string
}

func NewListV3OrganizationsByQuery(services *core.Services, urls map[string]string) *ListV3OrganizationsByQuery {
	return &ListV3OrganizationsByQuery{services: services, urls: urls}
}

func (route *ListV3OrganizationsByQuery) Register(e *echo.Echo) {
	e.GET("/v3/organizations", route.Handle)
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
