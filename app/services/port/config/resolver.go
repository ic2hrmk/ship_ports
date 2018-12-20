package config

import (
	"fmt"

	"github.com/ic2hrmk/ship_ports/shared/env"
)

//
// All env. vars that are used by service
//
const (
	mongoURL = "MONGO_URL"
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

	if config.MongoURL, err = env.GetStringVar(mongoURL); err != nil {
		return fmt.Errorf("failed to read env. var [%s]: %s", mongoURL, err)
	}

	return nil
}
