package errors

import (
	"log"

	"github.com/ic2hrmk/ship_ports/app/services/port/pb/port"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Internal(err error) error {
	log.Printf("[port-srv] Internal error: %s", err)
	return status.Error(codes.Code(port.PortDomainServiceErrorCode_Internal), err.Error())
}

func InvalidRequest(err error) error {
	log.Printf("[port-srv] Invalid request: %s", err)
	return status.Error(codes.Code(port.PortDomainServiceErrorCode_InvalidRequest), err.Error())
}
