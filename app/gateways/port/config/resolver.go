package config

import (
	"fmt"

	"github.com/ic2hrmk/ship_ports/shared/env"
)

//
// All env. vars that are used by service
//
const (
	portServiceAddress = "SERVICE_PORT_ADDRESS"

	portMaxImportFileSize = "PORT_MAX_IMPORT_FILE_SIZE"
)

//
// Reads all configuration from env, db etc.
//
func ResolveConfigurations() (*ConfigurationContainer, error) {
	c := &ConfigurationContainer{}

	if err := resolveEnvConfiguration(c); err != nil {
		return nil, err
	}

	return c, nil
}

//
// Reads all variables, stored in env. file
//
func resolveEnvConfiguration(config *ConfigurationContainer) error {
	var err error

	if config == nil {
		config = &ConfigurationContainer{}
	}

	if config.PortDomainServiceAddress, err = env.GetStringVar(portServiceAddress); err != nil {
		return fmt.Errorf("failed to read env. var [%s]: %s", portServiceAddress, err)
	}

	if config.PortMaxImportSizeBytes, err = env.GetInt64Var(portMaxImportFileSize); err != nil {
		return fmt.Errorf("failed to read env. var [%s]: %s", portMaxImportFileSize, err)
	}

	return nil
}
