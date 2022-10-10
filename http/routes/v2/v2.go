package v2

import (
	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
)

func Register(router *echo.Echo, services *core.Services, urls map[string]string) {
	NewGetInfo(services, urls).Register(router)
	NewGetAppByGuid(services, urls).Register(router)
	NewListServicePlansByQuery(services, urls).Register(router)
	NewListUsersByQuery(services, urls).Register(router)
}
