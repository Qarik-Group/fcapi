package routes

import (
	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"

	"codeberg.org/ess/fcapi/http/routes/fcapi"
	"codeberg.org/ess/fcapi/http/routes/general"
	"codeberg.org/ess/fcapi/http/routes/v2"
	"codeberg.org/ess/fcapi/http/routes/v3"
)

func Register(router *echo.Echo, services *core.Services, urls map[string]string) {
	fcapi.Register(router, services, urls)
	v2.Register(router, services, urls)
	v3.Register(router, services, urls)
	general.Register(router, services, urls)

	//// System Routes
	//NewGetInfo(services, urls).Register(router)

	//// App Routes
	//NewGetAppByGuid(services, urls).Register(router)

	//// Task Routes
	//NewCreateTask(services, urls).Register(router)
	//NewGetTaskByGuid(services, urls).Register(router)

	//// User Routes
	//NewListUsersByQuery(services, urls).Register(router)

	//// V3 Role Routes
	//NewListV3RolesByQuery(services, urls).Register(router)
	//NewCreateV3Role(services, urls).Register(router)

	//// V3 Organization Routes
	//NewListV3OrganizationsByQuery(services, urls).Register(router)

	//// Service Plan Routes
	//NewListServicePlansByQuery(services, urls).Register(router)

	//// V3 Space Routes
	//NewListV3SpacesByQuery(services, urls).Register(router)
	//NewCreateV3Space(services, urls).Register(router)

	//// Service Broker Routes
	//NewCreateServiceBroker(services, urls).Register(router)

	//// CAPI Metadata Routes
	//NewGetRoot(services, urls).Register(router)
}
