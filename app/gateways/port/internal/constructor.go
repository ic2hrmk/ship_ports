package internal

import (
	"github.com/emicklei/go-restful"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/ic2hrmk/ship_ports/app"
)

func NewPortDomainService(
	config *portDomainGatewayConfig,
) app.MicroService {
	service := &portDomainGateway{
		config:       config,
		webContainer: restful.NewContainer(),
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
