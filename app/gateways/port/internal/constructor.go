package internal

import (
	portPb "github.com/ic2hrmk/ship_ports/app/services/port/pb/port"

	"github.com/emicklei/go-restful"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/ic2hrmk/ship_ports/app"
)

//
// Port Gateway micro-service
//
type portDomainGateway struct {
	webContainer *restful.Container
	config       *portDomainGatewayConfig

	//
	// Clients to back-services
	//
	portServiceClient portPb.PortDomainServiceClient
}

func NewPortDomainService(
	config *portDomainGatewayConfig,
	portServiceClient portPb.PortDomainServiceClient,
) app.MicroService {
	service := &portDomainGateway{
		config:            config,
		webContainer:      restful.NewContainer(),
		portServiceClient: portServiceClient,
	}

	service.init()

	return service
}

//
// =============== Additional configuration section ================
//

//
// Configuration container
//
type portDomainGatewayConfig struct {
	MaxImportFileSize int64
}

//
// Configuration container builder
//
type portDomainGatewayConfigBuilder struct {
	maxImportFileSize int64
}

func NewPortGatewayConfigBuilder() *portDomainGatewayConfigBuilder {
	return &portDomainGatewayConfigBuilder{}
}

//
// Builds gateway configuration object. It's the only way to initialize it's settings
//
func (rcv *portDomainGatewayConfigBuilder) Build() (*portDomainGatewayConfig, error) {
	if err := rcv.Validate(); err != nil {
		return nil, err
	}

	return &portDomainGatewayConfig{
		MaxImportFileSize: rcv.maxImportFileSize,
	}, nil
}

//
// Validates acquired settings
//
func (rcv *portDomainGatewayConfigBuilder) Validate() error {
	return validation.ValidateStruct(rcv,
		validation.Field(&rcv.maxImportFileSize, validation.Required),
	)
}

//
// ============ Configuration setters ================
//

func (rcv *portDomainGatewayConfigBuilder) SetImportFileMaxSize(size int64) (*portDomainGatewayConfigBuilder) {
	rcv.maxImportFileSize = size
	return rcv
}
