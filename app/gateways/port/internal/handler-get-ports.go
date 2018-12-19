package internal

import (
	"log"

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

	log.Println(limit, offset)

	out := &representation.PortListResponse{

	}

	/*

	//
	// Request information
	//
	portsDetails, err := rcv.portServiceClient.GetPorts(request.Request.Context(), &portPb.GetPortsRequest{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		helpers.ResponseWithInternalError(response, err)
	}

	//
	// Assemble response
	//
	out.Items = make([]*representation.PortEntityResponse, len(portsDetails.GetItems()))
	out.Found = len(portsDetails.GetItems())

	for i, port := range portsDetails.GetItems() {
		out[i] = &representation.PortEntityResponse{
			//...
		}
	}

	*/

	helpers.ResponseWithOK(response, out)
}
