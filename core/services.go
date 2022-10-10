package core

import (
	cf "github.com/cloudfoundry-community/go-cfclient"
)

type Services struct {
	Brokers BrokerService
	Orgs    OrganizationService
	Spaces  SpaceService
	Logger  Logger
}

type BrokerService interface {
	Add(string, string, string, string, string) (cf.ServiceBroker, error)
	ByName(string) (cf.ServiceBroker, error)
	Reset()
}

type OrganizationService interface {
	Get(string) (cf.V3Organization, error)
	All() []cf.V3Organization
	Reset()
}

type SpaceService interface {
	Add(cf.V3Organization, string) (cf.V3Space, error)
	ByOrg(cf.V3Organization) []cf.V3Space
	ByOrgAndName(cf.V3Organization, string) (cf.V3Space, error)
	Reset()
}

func (services *Services) Reset() {
	services.Orgs.Reset()
	services.Spaces.Reset()
	services.Brokers.Reset()
}
