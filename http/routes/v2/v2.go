package v2

import (
	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
)

func Register(router *echo.Echo, services *core.Services, urls map[string]string) {
	// System Routes
	NewGetInfo(services, urls).Register(router)

	// App Routes
	NewGetAppByGuid(services, urls).Register(router)

	// Service Plan Routes
	NewListServicePlansByQuery(services, urls).Register(router)

	// User Routes
	NewListUsersByQuery(services, urls).Register(router)
}
