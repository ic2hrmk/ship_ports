package internal

import (
	"github.com/ic2hrmk/ship_ports/app/services/port/errors"
	"github.com/ic2hrmk/ship_ports/app/services/port/pb/port"
	"golang.org/x/net/context"
)

//
// Finds all existing ports
//
func (rcv *portDomainService) FindAllPorts(
	ctx context.Context, in *port.FindAllPortsRequest,
) (*port.FindAllPortsResponse, error) {
	//
	// Validation
	//
	if err := in.Validate(); err != nil {
		return nil, errors.InvalidRequest(err)
	}

	//
	// Request handing
	//
	records, err := rcv.portRepository.FindAll(in.GetLimit(), in.GetOffset())
	if err != nil {
		return nil, errors.Internal(err)
	}

	//
	// Assemble response
	//
	out := &port.FindAllPortsResponse{
		Items: make([]*port.PortEntity, len(records)),
	}

	for i := range records {
		out.Items[i] = &port.PortEntity{
			PortID:      records[i].PortID,
			Name:        records[i].Name,
			Code:        records[i].Code,
			Alias:       records[i].Alias,
			Unlocs:      records[i].Unlocs,
			Country:     records[i].Country,
			Regions:     records[i].Regions,
			Province:    records[i].Province,
			City:        records[i].City,
			Coordinates: records[i].Coordinates,
			Timezone:    records[i].Timezone,
		}
	}

	return out, nil
}
