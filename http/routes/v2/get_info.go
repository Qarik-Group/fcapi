package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fcapi/core"
	"github.com/starkandwayne/fcapi/http/routes/registry"
)

type GetInfo struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewGetInfo(services *core.Services, urls map[string]string) *GetInfo {
	return &GetInfo{
		path:     "/v1/info",
		verb:     registry.GET,
		ready:    true,
		services: services,
		urls:     urls,
	}
}

func (route *GetInfo) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *GetInfo) Path() string {
	return route.path
}

func (route *GetInfo) Ready() bool {
	return route.ready
}

func (route *GetInfo) Verb() registry.Verb {
	return route.verb
}

func (route *GetInfo) Handle(c echo.Context) error {
	route.services.Logger(c)

	uaaURL, ok := route.urls["uaaURL"]
	if !ok {
		uaaURL = "https://uaa.example.com"
	}

	cfURL, ok := route.urls["cfURL"]
	if !ok {
		cfURL = "https://cf.example.com"
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"authorization_endpoint":       uaaURL,
			"token_endpoint":               uaaURL,
			"logging_endpoint":             cfURL,
			"name":                         "",
			"build":                        "",
			"support":                      "https://support.example.com",
			"version":                      0,
			"description":                  "",
			"min_cli_version":              "6.23.0",
			"min_recommended_cli_version":  "6.23.0",
			"api_version":                  "2.103.0",
			"app_ssh_endpoint":             "ssh.example.com:2222",
			"app_ssh_host_key_fingerprint": "00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:01",
			"app_ssh_oauth_client":         "ssh-proxy",
			"doppler_logging_endpoint":     "wss://doppler.example.com:443",
			"routing_endpoint":             "https://api.example.com/routing",
		},
	)
}
