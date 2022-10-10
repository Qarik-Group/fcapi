package memory

import (
	"errors"
	"sync"

	cf "github.com/cloudfoundry-community/go-cfclient"

	"codeberg.org/ess/fcapi/core"
)

type SpaceService struct {
	spaces map[string][]cf.V3Space
	locker sync.Mutex
}

func NewSpaceService() *SpaceService {
	service := &SpaceService{}
	service.Reset()

	return service
}

func (service *SpaceService) Add(org cf.V3Organization, name string) (cf.V3Space, error) {
	service.locker.Lock()
	defer service.locker.Unlock()

	return service.add(org, name)
}

func (service *SpaceService) add(org cf.V3Organization, name string) (cf.V3Space, error) {
	_, err := service.byOrgAndName(org, name)
	if err == nil {
		return cf.V3Space{}, errors.New("space already exists")
	}

	spaceGUID, _ := core.GenGUID()
	space := cf.V3Space{
		GUID: spaceGUID,
		Name: name,
	}

	service.spaces[org.GUID] = append(service.spaces[org.GUID], space)

	return space, nil
}

func (service *SpaceService) ByOrgAndName(org cf.V3Organization, name string) (cf.V3Space, error) {
	service.locker.Lock()
	defer service.locker.Unlock()

	return service.byOrgAndName(org, name)
}

func (service *SpaceService) byOrgAndName(org cf.V3Organization, name string) (cf.V3Space, error) {
	for _, space := range service.byOrg(org) {
		if space.Name == name {
			return space, nil
		}
	}

	return cf.V3Space{}, errors.New("no such space")
}

func (service *SpaceService) ByOrg(org cf.V3Organization) []cf.V3Space {
	service.locker.Lock()
	defer service.locker.Unlock()

	return service.byOrg(org)
}

func (service *SpaceService) byOrg(org cf.V3Organization) []cf.V3Space {
	if service.spaces[org.GUID] == nil {
		service.spaces[org.GUID] = make([]cf.V3Space, 0)
	}

	return service.spaces[org.GUID]
}

//func (service *SpaceService) All() []cf.V3Organization {
//service.locker.Lock()
//defer service.locker.Unlock()

//return service.all()
//}

//func (service *SpaceService) all() []cf.V3Organization {
//spaces := make([]cf.V3Space, 0)
//for _, space := range service.spaces {
//spaces = append(spaces, space)
//}

//return spaces
//}

func (service *SpaceService) Reset() {
	service.locker.Lock()
	defer service.locker.Unlock()

	service.reset()
}

func (service *SpaceService) reset() {
	service.spaces = make(map[string][]cf.V3Space)
}
