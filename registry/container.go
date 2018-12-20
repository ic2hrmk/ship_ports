package registry

import (
	"fmt"

	"github.com/ic2hrmk/ship_ports/app"
)

type registryContainer map[string]FactoryMethod

func NewRegistryContainer() registryContainer {
	return make(registryContainer)
}

type FactoryMethod func() (app.MicroService, error)

func (c *registryContainer) Add(serviceName string, fabric FactoryMethod) {
	(*c)[serviceName] = fabric
}

func (c *registryContainer) Get(serviceName string) (FactoryMethod, error) {
	entry, ok := (*c)[serviceName]
	if !ok {
		return nil, fmt.Errorf("service [%s] isn't registered", serviceName)
	}

	return entry, nil
}
