package internal

import (
	"log"

	"github.com/ic2hrmk/ship_ports/app/services/port/errors"
	"github.com/ic2hrmk/ship_ports/app/services/port/pb/port"
	"golang.org/x/net/context"
)

func (rcv *portDomainService) GetPorts(
	ctx context.Context, in *port.GetPortsRequest,
) (*port.GetPortsResponse, error) {
	//
	// Validation
	//
	if err := in.Validate(); err != nil {
		return nil, errors.InvalidRequest(err)
	}

	//
	// Request handing
	//

	log.Println("GET PORT")

	//
	// Assemble response
	//
	out := &port.GetPortsResponse{}

	return out, nil
}
