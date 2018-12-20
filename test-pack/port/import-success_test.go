package port

import (
	"fmt"
	"github.com/ic2hrmk/ship_ports/shared/gateway/representation"
	"testing"

	"github.com/go-resty/resty"
	"github.com/ic2hrmk/ship_ports/shared/env"
)

//
// Looks how service reacts on case, when file is OK
//
func TestImport_Success(t *testing.T) {
	//
	// Get file fixture path
	//
	testFilePath, err := env.GetStringVar(correctFilePath)
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

	if !resp.IsSuccess() {
		t.Fatalf("wrong server response - errors are detected, %v", resp)
	}

	if errResponse != nil && errResponse.Message != "" {
		t.Fatalf("error response is not nil, %v", *errResponse)
	}
}
