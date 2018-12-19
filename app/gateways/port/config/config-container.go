package config

import (
	"github.com/go-ozzo/ozzo-validation"
)

//
// All available configurations for the micro-service
//
type ConfigurationContainer struct {
	//
	// Addresses of all connected services
	//
	PortDomainServiceAddress string

	//
	// Service's custom settings
	//
	PortMaxImportSizeBytes int64
}

func (c *ConfigurationContainer) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.PortDomainServiceAddress, validation.Required),
		validation.Field(&c.PortMaxImportSizeBytes, validation.Required),
	)
}
