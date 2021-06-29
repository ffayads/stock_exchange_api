package utils

import (
	"strconv"

	httpmodels "github.com/api/stock_exchange_api/httpmodels"
)

func BindDelete(id string) *httpmodels.DeleteRequest {
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil
	}
	bind := &httpmodels.DeleteRequest{
		ID: uint(u64),
	}
	return bind
}
