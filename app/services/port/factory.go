package port

import (
	sharedMongo "github.com/ic2hrmk/ship_ports/shared/persistence/mongo"

	"github.com/globalsign/mgo"
	"github.com/ic2hrmk/ship_ports/app"
	"github.com/ic2hrmk/ship_ports/app/services/port/config"
	"github.com/ic2hrmk/ship_ports/app/services/port/internal"
	"github.com/ic2hrmk/ship_ports/app/services/port/persistence/repository/mongo"
)

const ServiceName = "port-srv"

//
// Service constructor
//
func FactoryMethod() (app.MicroService, error) {
	//
	// Resolve configurations
	//	- service's configurations
	//
	configurations, err := resolveConfigurations()
	if err != nil {
		return nil, err
	}

	//
	// Init. persistence
	//
	mongoDB, err := initMongoDB(configurations.MongoURL, ServiceName)
	if err != nil {
		return nil, err
	}

	// Init. port repository
	portRepository := mongo.NewPortRepository(mongoDB)

	return internal.NewPortDomainService(
		portRepository,
	), nil
}

func resolveConfigurations() (*config.ConfigurationContainer, error) {
	return config.ResolveConfigurations()
}

func initMongoDB(mongoURL, dbName string) (*mgo.Database, error) {
	return sharedMongo.InitConnection(mongoURL, dbName)
}
