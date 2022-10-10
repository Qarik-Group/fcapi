package v3

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
)

type CreateV3Space struct {
	services *core.Services
	urls     map[string]string
}

func NewCreateV3Space(services *core.Services, urls map[string]string) *CreateV3Space {
	return &CreateV3Space{services: services, urls: urls}
}

func (route *CreateV3Space) Register(e *echo.Echo) {
	e.POST("/v3/spaces", route.Handle)
}

func (route *CreateV3Space) Handle(c echo.Context) error {
	route.services.Logger(c)

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "could not read payload")
	}

	data := make(map[string]interface{})
	if err = json.Unmarshal(body, &data); err != nil {
		return c.JSON(http.StatusInternalServerError, "could not decode payload: "+err.Error())
	}

	name := data["name"].(string)
	if len(name) == 0 {
		return c.JSON(http.StatusInternalServerError, "name cannot be blank")
	}

	orgGUID := data["relationships"].(map[string]interface{})["organization"].(map[string]interface{})["data"].(map[string]interface{})["guid"].(string)

	if len(orgGUID) == 0 {
		return c.JSON(http.StatusInternalServerError, "org guid cannot be blank")
	}

	org, err := route.services.Orgs.Get(orgGUID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "no such org")
	}

	_, err = route.services.Spaces.ByOrgAndName(org, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "name already taken")
	}

	space, err := route.services.Spaces.Add(org, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(
		http.StatusCreated,
		space,
	)
}
