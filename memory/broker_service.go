package memory

import (
	"errors"
	"sync"
	"time"

	cf "github.com/cloudfoundry-community/go-cfclient"

	"github.com/starkandwayne/fcapi/core"
)

type BrokerService struct {
	brokers []cf.ServiceBroker
	locker  sync.Mutex
}

func NewBrokerService() *BrokerService {
	service := &BrokerService{}
	service.Reset()

	return service
}

func (service *BrokerService) Add(name string, username string, password string, url string, spaceGUID string) (cf.ServiceBroker, error) {
	service.locker.Lock()
	defer service.locker.Unlock()

	return service.add(name, username, password, url, spaceGUID)
}

func (service *BrokerService) add(name string, username string, password string, url string, spaceGUID string) (cf.ServiceBroker, error) {
	_, err := service.byName(name)
	if err == nil {
		return cf.ServiceBroker{}, errors.New("name already taken")
	}

	brokerGUID, _ := core.GenGUID()
	now := time.Now().String()

	broker := cf.ServiceBroker{
		Guid:      brokerGUID,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
		BrokerURL: url,
		Username:  username,
		Password:  password,
		SpaceGUID: spaceGUID,
	}

	service.brokers = append(service.brokers, broker)

	return broker, nil
}

func (service *BrokerService) ByName(name string) (cf.ServiceBroker, error) {
	service.locker.Lock()
	defer service.locker.Unlock()

	return service.byName(name)
}

func (service *BrokerService) byName(name string) (cf.ServiceBroker, error) {
	for _, broker := range service.brokers {
		if broker.Name == name {
			return broker, nil
		}
	}

	return cf.ServiceBroker{}, errors.New("broker not found")
}

func (service *BrokerService) Reset() {
	service.locker.Lock()
	defer service.locker.Unlock()

	service.reset()
}

func (service *BrokerService) reset() {
	service.brokers = make([]cf.ServiceBroker, 0)
}
