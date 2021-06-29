package helpers

import (
	"github.com/api/stock_exchange_api/httpmodels"
	"github.com/api/stock_exchange_api/models"
	"github.com/api/stock_exchange_api/repository"
)

func GetRecords(values *httpmodels.Filter, user *models.Users) []*httpmodels.Record {
	return repository.RecordsConvertToListResponse(repository.GetRecordsDetail(values, user))
}

func GetRecord(values *httpmodels.GetRecordDetail, user *models.Users) *httpmodels.GetRecordDetailResponse {
	first := repository.GetRecordsFirst(values, user)
	last := repository.GetRecordsLast(values, user)
	return repository.RecordDetailConvertToResponse(values.InvestmentID, first.CurrentValue, last.TotalValue)
}
