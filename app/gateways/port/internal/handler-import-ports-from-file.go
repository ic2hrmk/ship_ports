package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/emicklei/go-restful"
	"github.com/ic2hrmk/ship_ports/app/gateways/port/errors"
	"github.com/ic2hrmk/ship_ports/app/gateways/port/representation"
	"github.com/ic2hrmk/ship_ports/app/services/port/pb/port"
	"github.com/ic2hrmk/ship_ports/shared/gateway/helpers"
)

func (rcv *portDomainGateway) importPortsFromFile(
	request *restful.Request,
	response *restful.Response,
) {
	var err error

	// TODO: enable file size limit
	//request.Request.Body = http.MaxBytesReader(response, request.Request.Body, rcv.config.MaxImportFileSize)

	if err := request.Request.ParseMultipartForm(rcv.config.MaxImportFileSize); err != nil {
		helpers.ResponseWithBadRequest(response, err, errors.ErrFileIsTooBig)
		return
	}

	file, _, err := request.Request.FormFile(fileParameterName)
	if err != nil {
		helpers.ResponseWithBadRequest(response, err, errors.ErrFileNotAttached)
		return
	}

	//
	// Try to read file as a stream
	//
	parsedEntities := make(chan *port.PortEntity)
	done := make(chan struct{})

	//
	// Sync.
	//
	wg := sync.WaitGroup{}

	//
	// Add producer
	//
	wg.Add(1)

	go func() {
		err = rcv.decodePortEntityStream(file, parsedEntities, done)
		wg.Done()
	}()

	//
	// Attach consumer
	//
	wg.Add(1)

	go func() {
		rcv.upsertPortEntityFromStream(parsedEntities, done)
		wg.Done()
	}()

	wg.Wait()

	if err != nil {
		helpers.ResponseWithBadRequest(response, err, errors.ErrInconsistentJson)
		return
	}

	helpers.ResponseWithNoContent(response)
}

func (rcv *portDomainGateway) decodePortEntityStream(
	r io.Reader,
	parsed chan *port.PortEntity,
	done chan struct{},
) error {
	endSignal := struct{}{}

	defer func() {
		done <- endSignal
	}()

	dec := json.NewDecoder(r)

	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch t.(type) {
		case json.Delim:
			// Do nothing, skip delimeter
		default:
			if dec.More() {
				var portID string
				var isPortID bool

				if portID, isPortID = t.(string); !isPortID {
					return fmt.Errorf("inconsistent JSON")
				}

				row := &representation.ImportedPortEntity{}

				if err := dec.Decode(&row); err != nil {
					return fmt.Errorf("failed to decodeentity: %s", err)
				}

				parsed <- &port.PortEntity{
					PortID:      portID,
					Name:        row.Name,
					Coordinates: row.Coordinates,
					City:        row.City,
					Province:    row.Province,
					Country:     row.Country,
					Alias:       row.Alias,
					Regions:     row.Regions,
					Timezone:    row.Timezone,
					Unlocs:      row.Unlocs,
					Code:        row.Code,
				}
			}
		}
	}

	return nil
}

func (rcv *portDomainGateway) upsertPortEntityFromStream(parsed chan *port.PortEntity, done chan struct{}) {
	for {
		select {
		case record := <-parsed:
			_, err := rcv.portServiceClient.SavePort(context.Background(), &port.SavePortRequest{Port: record})
			if err != nil {
				log.Printf("  --->  | Failed to persist entity [PortID=%s], %s", record.PortID, err)
			}

		case <-done:
			return
		}
	}
}
