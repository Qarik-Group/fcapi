package v3

import (
	"net/http"
	"strings"

	cf "github.com/cloudfoundry-community/go-cfclient"
	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
	"codeberg.org/ess/fcapi/http/responses"
	"codeberg.org/ess/fcapi/http/routes/registry"
)

type ListV3SpacesByQuery struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewListV3SpacesByQuery(services *core.Services, urls map[string]string) *ListV3SpacesByQuery {
	return &ListV3SpacesByQuery{
		path:     "/v3/roles",
		verb:     registry.GET,
		ready:    true,
		services: services,
		urls:     urls,
	}
}

func (route *ListV3SpacesByQuery) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *ListV3SpacesByQuery) Path() string {
	return route.path
}

func (route *ListV3SpacesByQuery) Ready() bool {
	return route.ready
}

func (route *ListV3SpacesByQuery) Verb() registry.Verb {
	return route.verb
}

func (route *ListV3SpacesByQuery) Handle(c echo.Context) error {
	route.services.Logger(c)

	orgGUID := c.QueryParam("organization_guids")
	org, err := route.services.Orgs.Get(orgGUID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "no such org")
	}

	spaces := route.services.Spaces.ByOrg(org)

	names := c.QueryParam("names")
	if len(names) > 0 {
		filtered := make([]cf.V3Space, 0)

		for _, name := range strings.Split(names, ",") {
			for _, space := range spaces {
				if space.Name == name {
					filtered = append(filtered, space)
				}
			}
		}

		spaces = filtered
	}

	output := &responses.SpaceCollection{
		Resources: spaces,
		Pagination: &responses.Pagination{
			TotalPages:   1,
			TotalResults: len(spaces),
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
