package port

import (
	"github.com/ic2hrmk/ship_ports/app"
	"github.com/ic2hrmk/ship_ports/app/services/port/internal"
)

const ServiceName = "port-srv"

func FactoryMethod() (app.MicroService, error) {
	//mongoDB, err := initMongoDB()
	//if err != nil {
	//	return nil, err
	//}
	//
	//portRepository := mongo.NewPortRepository(mongoDB)

	return internal.NewPortDomainService(), nil
}
