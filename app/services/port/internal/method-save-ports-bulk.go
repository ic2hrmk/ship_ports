package internal

import (
	"github.com/ic2hrmk/ship_ports/app/services/port/errors"
	"github.com/ic2hrmk/ship_ports/app/services/port/pb/port"
	"github.com/ic2hrmk/ship_ports/app/services/port/persistence/model"
	"golang.org/x/net/context"
)

//
// Saves ports to persistence via bulk
//
func (rcv *portDomainService) SavePortsBulk(
	ctx context.Context, in *port.SavePortsBulkRequest,
) (*port.SavePortsBulkResponse, error) {
	//
	// Validation
	//
	if err := in.Validate(); err != nil {
		return nil, errors.InvalidRequest(err)
	}

	//
	// Discard empty bulk inserts
	//
	if len(in.GetItems()) == 0 {
		return &port.SavePortsBulkResponse{}, nil
	}

	//
	// Request handing
	//
	records := make([]*model.Port, len(in.GetItems()))

	for i, portDetails := range in.GetItems() {
		records[i] = &model.Port{
			PortID:      portDetails.GetPortID(),
			Name:        portDetails.GetName(),
			Code:        portDetails.GetCode(),
			Alias:       portDetails.GetAlias(),
			Unlocs:      portDetails.GetUnlocs(),
			Country:     portDetails.GetCountry(),
			Regions:     portDetails.GetRegions(),
			Province:    portDetails.GetProvince(),
			City:        portDetails.GetCity(),
			Coordinates: portDetails.GetCoordinates(),
			Timezone:    portDetails.GetTimezone(),
		}
	}

	if err := rcv.portRepository.SaveBulk(records); err != nil {
		return nil, errors.Internal(err)
	}

	//
	// Assemble response
	//
	return &port.SavePortsBulkResponse{}, nil
}
