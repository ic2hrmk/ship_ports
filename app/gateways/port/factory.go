package port

import (
	"github.com/ic2hrmk/ship_ports/app"
	"github.com/ic2hrmk/ship_ports/app/gateways/port/config"
	"github.com/ic2hrmk/ship_ports/app/gateways/port/internal"
)

const ServiceName = "port-gtw"

func FactoryMethod() (app.MicroService, error) {
	//
	// Resolve configurations
	//	- clients's configurations
	//  - gateway's configurations
	//
	configurations, err := resolveConfigurations()
	if err != nil {
		return nil, err
	}

	//
	// Init. gateway configurations
	//
	gatewayConfigurationBuilder := internal.NewPortGatewayConfigBuilder()

	gatewayConfigurationBuilder.SetImportFileMaxSize(configurations.PortMaxImportSizeBytes)
	// ... set any other configurations here

	gatewayConfiguration, err := gatewayConfigurationBuilder.Build()
	if err != nil {
		return nil, err
	}

	//
	// Init. clients
	//

	return internal.NewPortDomainService(gatewayConfiguration), nil
}

func resolveConfigurations() (*config.ConfigurationContainer, error) {
	return config.ResolveConfigurations()
}
