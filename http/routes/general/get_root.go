package general

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"codeberg.org/ess/fcapi/core"
	"codeberg.org/ess/fcapi/http/responses"
	"codeberg.org/ess/fcapi/http/routes/registry"
)

type GetRoot struct {
	path     string
	verb     registry.Verb
	ready    bool
	services *core.Services
	urls     map[string]string
}

func NewGetRoot(services *core.Services, urls map[string]string) *GetRoot {
	return &GetRoot{
		path:     "/",
		verb:     registry.GET,
		ready:    true,
		services: services,
		urls:     urls,
	}
}

func (route *GetRoot) Register(router *echo.Echo) {
	registry.Register(router, route)
}

func (route *GetRoot) Path() string {
	return route.path
}

func (route *GetRoot) Ready() bool {
	return route.ready
}

func (route *GetRoot) Verb() registry.Verb {
	return route.verb
}

func (route *GetRoot) Handle(c echo.Context) error {
	route.services.Logger(c)

	uaaURL, ok := route.urls["uaaURL"]
	if !ok {
		uaaURL = "https://uaa.example.com"
	}

	cfURL, ok := route.urls["cfURL"]
	if !ok {
		cfURL = "https://cf.example.com"
	}

	output := &responses.RootResp{
		Links: &responses.RootLinkCollection{
			Self: &responses.Rootref{Href: cfURL},

			CloudControllerV2: &responses.Rootref{
				Href: cfURL + "/v2",
				Meta: map[string]string{
					"version": "2.155.0",
				},
			},

			CloudControllerV3: &responses.Rootref{
				Href: cfURL + "/v3",
				Meta: map[string]string{
					"version": "3.99.0",
				},
			},

			NetworkPolicyV0: &responses.Rootref{Href: cfURL + "/networking/v0/external"},
			NetworkPolicyV1: &responses.Rootref{Href: cfURL + "/networking/v1/external"},
			Login:           &responses.Rootref{Href: uaaURL},
			UAA:             &responses.Rootref{Href: uaaURL},
			Routing:         &responses.Rootref{Href: cfURL + "/routing"},
			Logging:         &responses.Rootref{Href: cfURL + "/jelogging"},
			LogCache:        &responses.Rootref{Href: cfURL + "/jelogcache"},
			LogStream:       &responses.Rootref{Href: cfURL + "/jelogstream"},

			AppSSH: &responses.Rootref{
				Href: "ssh.example.org:2222",
				Meta: map[string]string{
					"host_key_fingerprint": "Y411oivJwZCUQnXHq83mdM5SKCK4ftyoSXI31RRe4Zs",
					"oauth_client":         "ssh-proxy",
				},
			},
		},
	}

	return c.JSON(
		http.StatusOK,
		output,
	)
}
