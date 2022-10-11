package v3

import (
	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fcapi/core"
)

func Register(router *echo.Echo, services *core.Services, urls map[string]string) {
	NewCreateServiceBroker(services, urls).Register(router)
	NewCreateTask(services, urls).Register(router)
	NewCreateV3Role(services, urls).Register(router)
	NewCreateV3Space(services, urls).Register(router)
	NewGetTaskByGuid(services, urls).Register(router)
	NewListV3OrganizationsByQuery(services, urls).Register(router)
	NewListV3RolesByQuery(services, urls).Register(router)
	NewListV3SpacesByQuery(services, urls).Register(router)
}
