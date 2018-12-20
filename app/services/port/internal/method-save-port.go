package internal

import (
	"github.com/ic2hrmk/ship_ports/app/services/port/errors"
	"github.com/ic2hrmk/ship_ports/app/services/port/pb/port"
	"github.com/ic2hrmk/ship_ports/app/services/port/persistence/model"
	"golang.org/x/net/context"
)

//
// Saves port to persistence
//
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
	if _, err := rcv.portRepository.Save(&model.Port{
		PortID:      in.GetPort().GetPortID(),
		Name:        in.GetPort().GetName(),
		Code:        in.GetPort().GetCode(),
		Alias:       in.GetPort().GetAlias(),
		Unlocs:      in.GetPort().GetUnlocs(),
		Country:     in.GetPort().GetCountry(),
		Regions:     in.GetPort().GetRegions(),
		Province:    in.GetPort().GetProvince(),
		City:        in.GetPort().GetCity(),
		Coordinates: in.GetPort().GetCoordinates(),
		Timezone:    in.GetPort().GetTimezone(),
	}); err != nil {
		return nil, errors.Internal(err)
	}

	//
	// Assemble response
	//
	return &port.SavePortResponse{}, nil
}
