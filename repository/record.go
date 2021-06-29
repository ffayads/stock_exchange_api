package repository

import (
	"github.com/api/stock_exchange_api/config/db"
	"github.com/api/stock_exchange_api/httpmodels"
	"github.com/api/stock_exchange_api/models"
	"github.com/api/stock_exchange_api/utils"
)

func RecordsConvertToListResponse(users []models.Records) []*httpmodels.Record {
	usersResp := []*httpmodels.Record{}
	for _, user := range users {
		usersResp = append(usersResp, RecordConvertToResponse(&user))
	}
	return usersResp
}

func RecordConvertToResponse(user *models.Records) *httpmodels.Record {

	userRes := &httpmodels.Record{
		ID:           user.ID,
		Investment:   user.InvestmentsID,
		CurrentValue: user.CurrentValue,
		Yield:        user.Yield,
		TotalValue:   user.TotalValue,
		TypeRecords:  user.TypeRecords,
	}
	return userRes
}

/*func RecordsDetailConvertToListResponse(records []models.Records, user *models.Users, params *httpmodels.GetRecordDetail) []*httpmodels.GetRecordDetailResponse {
	recordsResp := []*httpmodels.GetRecordDetailResponse{}
	for _, record := range records {
		params.InvestmentID = &record.InvestmentsID
		first := GetRecordsFirst(params, user)
		last := GetRecordsLast(params, user)
		recordsResp = append(recordsResp, RecordDetailConvertToResponse(&record, first.CurrentValue, last.TotalValue))
	}
	return recordsResp
}*/

func RecordDetailConvertToResponse(id uint, initValue, totalValue float64) *httpmodels.GetRecordDetailResponse {

	userRes := &httpmodels.GetRecordDetailResponse{
		Investment: id,
		InitValue:  initValue,
		TotalValue: totalValue,
		Gain:       (totalValue - initValue),
	}
	return userRes
}

func GetRecords(params *httpmodels.Filter) []models.Records {
	users := []models.Records{}
	query := db.DB
	query = utils.Filter(query, params)
	query.Find(&users)
	return users
}

func GetRecordsFirst(params *httpmodels.GetRecordDetail, user *models.Users) *models.Records {
	records := &models.Records{}
	query := db.DB.Preload("Investments", "users_id = ?", user.ID).Where("created_at >= ? and investments_id = ?", params.MinDate, params.InvestmentID).Order("created_at asc")
	err := query.First(records).Error
	if err != nil {
		return nil
	}
	return records
}

func GetRecordsLast(params *httpmodels.GetRecordDetail, user *models.Users) *models.Records {
	records := &models.Records{}
	query := db.DB.Preload("Investments", "users_id = ?", user.ID).Where("created_at <= ? and investments_id = ?", params.MaxDate, params.InvestmentID).Order("created_at desc")
	err := query.First(records).Error
	if err != nil {
		return nil
	}
	return records
}

func GetRecordsDetail(params *httpmodels.Filter, user *models.Users) []models.Records {
	users := []models.Records{}
	query := db.DB.Preload("Investments", "users_id = ?", user.ID)
	query = utils.Filter(query, params)
	query.Find(&users)
	return users
}

func GetRecordByFilters(params *httpmodels.Filter) *models.Records {
	users := &models.Records{}
	query := db.DB
	query = utils.Filter(query, params)
	query.First(&users)
	return users
}

func GetRecordByID(id interface{}) *models.Records {

	user := models.Records{}
	query := db.DB.Where("id = ?", id)
	if query.First(&user).Error != nil {
		return nil
	}
	return &user
}
func CreateRecord(params *httpmodels.CreateInvestmentRequest, investments *models.Investments) (*models.Records, error) {
	u := &models.Records{
		InvestmentsID: investments.ID,
		CurrentValue:  investments.CurrentValue,
		Yield:         0,
		TotalValue:    investments.CurrentValue,
	}
	err := u.Create()
	return u, err
}

func UpdateInvestmentRecord(params *httpmodels.UpdateInvestmentRequest, investments *models.Investments, currentValue, totalValue float64) (*models.Records, error) {
	var typeRecord int
	if params.Shares != 0 {
		typeRecord = models.RECORD_SHARE
		params.Yield = 0
	} else {
		typeRecord = models.RECORD_YIELD
	}
	u := &models.Records{
		InvestmentsID: investments.ID,
		CurrentValue:  currentValue,
		Yield:         params.Yield,
		TotalValue:    totalValue,
		TypeRecords:   typeRecord,
	}
	err := u.Create()
	return u, err
}

func DeleteRecord(params *httpmodels.DeleteRequest) error {
	c := models.Records{}
	if db.DB.Where("investments_id = ?", params.ID).First(&c).Error != nil {
		return nil
	}
	err := c.Delete()
	return err
}
