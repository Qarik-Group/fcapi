package general

import (
	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fcapi/core"
)

func Register(router *echo.Echo, services *core.Services, urls map[string]string) {
	NewGetRoot(services, urls).Register(router)
}
