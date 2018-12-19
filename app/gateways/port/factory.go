package port

import (
	"fmt"

	portPb "github.com/ic2hrmk/ship_ports/app/services/port/pb/port"

	"github.com/ic2hrmk/ship_ports/app"
	"github.com/ic2hrmk/ship_ports/app/gateways/port/config"
	"github.com/ic2hrmk/ship_ports/app/gateways/port/internal"
	"google.golang.org/grpc"
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
	portServiceClient, err := initPortServiceClient(configurations.PortDomainServiceAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to init. port service client, %s", err)
	}

	return internal.NewPortDomainService(
		gatewayConfiguration,
		portServiceClient,
	), nil
}

func resolveConfigurations() (*config.ConfigurationContainer, error) {
	return config.ResolveConfigurations()
}

func initPortServiceClient(address string) (portPb.PortDomainServiceClient, error) {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("did not connect: %s", err)
	}

	return portPb.NewPortDomainServiceClient(conn), nil
}
