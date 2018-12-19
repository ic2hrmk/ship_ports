package internal

import (
	"log"

	"github.com/ic2hrmk/ship_ports/app/services/port/errors"
	"github.com/ic2hrmk/ship_ports/app/services/port/pb/port"
	"golang.org/x/net/context"
)

func (rcv *portDomainService) SavePort(
	ctx context.Context, in *port.SavePortRequest,
) (*port.SavePortResponse, error) {
	//
	// Validation
	//
	if err := in.Validate(); err != nil {
		return nil, errors.InvalidRequest(err)
	}

	//
	// Request handing
	//

	log.Println("SAVE PORT")

	//
	// Assemble response
	//
	out := &port.SavePortResponse{}

	return out, nil
}
