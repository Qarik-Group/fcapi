package v3

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fcapi/core"
	"github.com/starkandwayne/fcapi/http/routes/registry"
)

type CreateServiceBroker struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewCreateServiceBroker(services *core.Services, urls map[string]string) *CreateServiceBroker {
	return &CreateServiceBroker{
		path:     "/v3/service_brokers",
		verb:     registry.POST,
		ready:    true,
		services: services,
		urls:     urls,
	}
}

func (route *CreateServiceBroker) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *CreateServiceBroker) Path() string {
	return route.path
}

func (route *CreateServiceBroker) Verb() registry.Verb {
	return route.verb
}

func (route *CreateServiceBroker) Ready() bool {
	return route.ready
}

func (route *CreateServiceBroker) Handle(c echo.Context) error {
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
		return c.JSON(http.StatusInternalServerError, "name can't be blank")
	}

	url := data["url"].(string)
	if len(url) == 0 {
		return c.JSON(http.StatusInternalServerError, "url can't be blank")
	}

	credentials := data["authentication"].(map[string]interface{})["credentials"].(map[string]interface{})
	username := credentials["username"].(string)
	password := credentials["password"].(string)

	broker, err := route.services.Brokers.Add(name, username, password, url, "")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(
		http.StatusCreated,
		broker,
	)
}
