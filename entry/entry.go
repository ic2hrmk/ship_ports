package main

import (
	"log"

	portGateway "github.com/ic2hrmk/ship_ports/app/gateways/port"
	portService "github.com/ic2hrmk/ship_ports/app/services/port"

	"github.com/ic2hrmk/ship_ports/registry"
	"github.com/ic2hrmk/ship_ports/shared/cmd"
	"github.com/ic2hrmk/ship_ports/shared/env"
)

//go:generate go run entry.go --kind=port-gtw --address=:8080 --env=docker-compose.env
//go:generate go run entry.go --kind=port-srv

func main() {
	//
	// Load startup flags
	//
	flags := cmd.LoadFlags()

	//
	// Load env.
	//
	if flags.EnvFile != "" {
		err := env.LoadEnvFile(flags.EnvFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	//
	// Select service
	//
	reg := registry.NewRegistryContainer()

	reg.Add(portGateway.ServiceName, portGateway.FactoryMethod)
	reg.Add(portService.ServiceName, portService.FactoryMethod)

	serviceFactory, err := reg.Get(flags.Kind)
	if err != nil {
		log.Fatal(err)
	}

	//
	// Create service
	//
	service, err := serviceFactory()
	if err != nil {
		log.Fatal(err)
	}

	//
	// Run till the death comes
	//
	log.Printf("[%s] started serving on '%s'", flags.Kind, flags.Address)
	log.Fatal(service.Serve(flags.Address))
}
