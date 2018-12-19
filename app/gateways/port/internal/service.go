package internal

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/ic2hrmk/ship_ports/app/gateways/port/representation"
	"github.com/ic2hrmk/ship_ports/shared/gateway/filters"
)

//
// Port Gateway micro-service
//
type portDomainGateway struct {
	webContainer *restful.Container
	config       *portDomainGatewayConfig
}

func (rcv *portDomainGateway) init() {
	ws := &restful.WebService{}

	ws.Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/ports/import").
		To(rcv.importPortsFromFile).
		Operation("importPortsFromFile").
		Consumes("multipart/form-data").
		Param(
			ws.FormParameter("file", "Import").DataType("file"),
		).
		Writes(representation.ImportFromFileResponse{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.ImportFromFileResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), representation.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.ErrorResponse{}))

	ws.Route(ws.GET("/ports").
		To(rcv.getPorts).
		Operation("getPorts").
		Param(ws.QueryParameter(limitParameterName, "Limit").DataType("integer")).
		Param(ws.QueryParameter(offsetParameterName, "Offset").DataType("integer")).
		Writes(representation.PortListResponse{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.PortListResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), representation.ErrorResponse{}))

	ws.Filter(filters.LogRequest)

	rcv.webContainer.Add(ws)
}

func (rcv *portDomainGateway) Serve(address string) error {
	return http.ListenAndServe(address, rcv.webContainer)
}