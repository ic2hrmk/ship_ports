package internal

import "github.com/ic2hrmk/ship_ports/app"

type portDomainService struct {

}

func NewPortDomainService() app.MicroService {
	return &portDomainService{

	}
}

func (rcv *portDomainService) Serve(address string) error {
	return nil
}
