package port

import (
	"fmt"
	"github.com/ic2hrmk/ship_ports/app/gateways/port/errors"
	"testing"

	"github.com/go-resty/resty"
	"github.com/ic2hrmk/ship_ports/shared/env"
	"github.com/ic2hrmk/ship_ports/shared/gateway/representation"
)

//
// Looks how service reacts on case, when file is damaged
//
func TestImport_Inconsistent(t *testing.T) {
	//
	// Get file fixture path
	//
	testFilePath, err := env.GetStringVar(inconsistentFilePath)
	if err != nil {
		t.Fatalf("failed to get fixture file path: %s", err)
	}

	//
	// Get target URL
	//
	portGatewayURL, err := env.GetStringVar(portGatewayURL)
	if err != nil {
		t.Fatalf("failed to get target URL: %s", err)
	}

	//
	// Assemble request
	//
	errResponse := &representation.ErrorResponse{}

	resp, err := resty.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFile("file", testFilePath).
		SetError(&errResponse).
		Post(fmt.Sprintf("%s/api/ports/import", portGatewayURL))

	if err != nil {
		t.Fatalf("failed to send upload request, %s", err)
	}

	if resp.IsSuccess() {
		t.Fatalf("wrong server response - errors are not detected")
	}

	if errResponse == nil {
		t.Fatalf("error response is nil")
	}

	if errResponse.Message != errors.ErrInconsistentJson {
		t.Fatalf("wrong error message [%s], expected [%s]", errResponse.Message, errors.ErrInconsistentJson)
	}
}
