package internal

import (
	"context"
	portPb "github.com/ic2hrmk/ship_ports/app/services/port/pb/port"

	"github.com/emicklei/go-restful"
	"github.com/ic2hrmk/ship_ports/app/gateways/port/errors"
	"github.com/ic2hrmk/ship_ports/app/gateways/port/representation"
	"github.com/ic2hrmk/ship_ports/shared/gateway/helpers"
)

func (rcv *portDomainGateway) getPorts(
	request *restful.Request,
	response *restful.Response,
) {
	var limit, offset uint64
	var err error

	//
	// Get limit and offset
	//
	if limit, err = getLimitParameter(request); err != nil {
		helpers.ResponseWithBadRequest(response, err, errors.ErrInvalidLimitParameter)
		return
	}

	if offset, err = getOffsetParameter(request); err != nil {
		helpers.ResponseWithBadRequest(response, err, errors.ErrInvalidOffsetParameter)
		return
	}

	//
	// Request information
	//
	portsDetails, err := rcv.portServiceClient.FindAllPorts(context.Background(),
		&portPb.FindAllPortsRequest{
			Limit:  limit,
			Offset: offset,
		})

	if err != nil {
		helpers.ResponseWithInternalError(response, err)
		return
	}

	//
	// Assemble response
	//
	out := &representation.PortListResponse{
		Items: make([]*representation.PortEntityResponse, len(portsDetails.GetItems())),
		Found: len(portsDetails.GetItems()),
	}

	for i, port := range portsDetails.GetItems() {
		out.Items[i] = &representation.PortEntityResponse{
			PortID:      port.GetPortID(),
			Name:        port.GetName(),
			Code:        port.GetCode(),
			Alias:       port.GetAlias(),
			Unlocs:      port.GetUnlocs(),
			Country:     port.GetCountry(),
			Regions:     port.GetRegions(),
			Province:    port.GetProvince(),
			City:        port.GetCity(),
			Coordinates: port.GetCoordinates(),
			Timezone:    port.GetTimezone(),
		}
	}

	helpers.ResponseWithOK(response, out)
}
