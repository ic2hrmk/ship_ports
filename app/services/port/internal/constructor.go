package internal

import (
	"fmt"
	"net"

	portPb "github.com/ic2hrmk/ship_ports/app/services/port/pb/port"

	"github.com/ic2hrmk/ship_ports/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type portDomainService struct {
}

func NewPortDomainService() app.MicroService {
	return &portDomainService{

	}
}

func (rcv *portDomainService) Serve(address string) error {
	return rcv.serve(address)
}

func (rcv *portDomainService) serve(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to acquire address [%s]: %s", address, err)
	}

	s := grpc.NewServer()
	portPb.RegisterPortDomainServiceServer(s, rcv)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

