package memory

import (
	"errors"
	"sync"
	"time"

	cf "github.com/cloudfoundry-community/go-cfclient"

	"github.com/starkandwayne/fcapi/core"
)

type OrganizationService struct {
	orgs   []cf.V3Organization
	locker sync.Mutex
}

func NewOrganizationService() *OrganizationService {
	service := &OrganizationService{}
	service.Reset()

	return service
}

func (service *OrganizationService) Get(guid string) (cf.V3Organization, error) {
	service.locker.Lock()
	defer service.locker.Unlock()

	return service.get(guid)
}

func (service *OrganizationService) get(guid string) (cf.V3Organization, error) {
	for _, org := range service.all() {
		if org.GUID == guid {
			return org, nil
		}
	}

	return cf.V3Organization{}, errors.New("no such org")
}

func (service *OrganizationService) All() []cf.V3Organization {
	service.locker.Lock()
	defer service.locker.Unlock()

	return service.all()
}

func (service *OrganizationService) all() []cf.V3Organization {
	return service.orgs
}

func (service *OrganizationService) Reset() {
	service.locker.Lock()
	defer service.locker.Unlock()

	service.reset()
}

func (service *OrganizationService) reset() {
	systemGUID, _ := core.GenGUID()
	now := time.Now()

	service.orgs = []cf.V3Organization{
		cf.V3Organization{
			Name:      "system",
			GUID:      systemGUID,
			CreatedAt: now.String(),
			UpdatedAt: now.String(),
		},
	}
}
