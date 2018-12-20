package internal

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/ic2hrmk/ship_ports/app/gateways/port/errors"
	"github.com/ic2hrmk/ship_ports/shared/gateway/helpers"
)

func (rcv *portDomainGateway) importPortsFromFile(
	request *restful.Request,
	response *restful.Response,
) {
	//
	// Set maximum request size and parse request
	//
	request.Request.Body = http.MaxBytesReader(response, request.Request.Body, rcv.config.MaxImportFileSize)

	if err := request.Request.ParseMultipartForm(rcv.config.MaxImportFileSize); err != nil {
		helpers.ResponseWithBadRequest(response, err, errors.ErrFileIsTooBig)
		return
	}

	//
	// Try to read file as a stream
	//
}
