package core

import (
	"github.com/labstack/echo/v4"
)

type Logger func(c echo.Context)
