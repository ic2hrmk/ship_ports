package internal

import (
	"fmt"
	"strconv"

	"github.com/emicklei/go-restful"
)

const (
	limitParameterName  = "limit"
	offsetParameterName = "offset"
)

func getLimitParameter(request *restful.Request) (uint64, error) {
	var limit uint64

	limitParam := request.QueryParameter(limitParameterName)
	if limitParam == "" {
		return defaultLimitValue, nil
	}

	limit, err := strconv.ParseUint(limitParam, 10, 64)
	if err != nil {
		return limit, fmt.Errorf("failed to parse parameter [%s], %s", limitParameterName, err)
	}

	return limit, nil
}

func getOffsetParameter(request *restful.Request) (uint64, error) {
	var limit uint64

	limitParam := request.QueryParameter(offsetParameterName)
	if limitParam == "" {
		return defaultOffsetValue, nil
	}

	limit, err := strconv.ParseUint(limitParam, 10, 64)
	if err != nil {
		return limit, fmt.Errorf("failed to parse parameter [%s], %s", offsetParameterName, err)
	}

	return limit, nil
}
