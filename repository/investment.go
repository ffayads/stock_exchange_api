package repository

import (
	"errors"
	"strconv"

	"github.com/api/stock_exchange_api/config/db"
	"github.com/api/stock_exchange_api/httpmodels"
	"github.com/api/stock_exchange_api/models"
	"github.com/api/stock_exchange_api/utils"
)

func InvestmentsConvertToListResponse(users []models.Investments) []*httpmodels.Investment {
	usersResp := []*httpmodels.Investment{}
	for _, user := range users {
		usersResp = append(usersResp, InvestmentConvertToResponse(&user))
	}
	return usersResp
}

func InvestmentConvertToResponse(user *models.Investments) *httpmodels.Investment {

	userRes := &httpmodels.Investment{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		User:         UserConvertToResponse(&user.Users),
		Company:      CompanyConvertToResponse(&user.Companies),
		Instrument:   InstrumentConvertToResponse(&user.Instruments),
		Shares:       user.Shares,
		InitValue:    user.InitValue,
		CurrentValue: user.CurrentValue,
		Status:       user.Status,
		Description:  user.Description,
		Records:      RecordsConvertToListResponse(user.Records),
	}
	return userRes
}

func InvestmentsDetailConvertToListResponse(users []models.Investments) []*httpmodels.Investment {
	usersResp := []*httpmodels.Investment{}
	for _, user := range users {
		usersResp = append(usersResp, InvestmentDetailConvertToResponse(&user))
	}
	return usersResp
}

func InvestmentDetailConvertToResponse(user *models.Investments) *httpmodels.Investment {

	userRes := &httpmodels.Investment{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Instrument:   InstrumentConvertToResponse(&user.Instruments),
		Shares:       user.Shares,
		InitValue:    user.InitValue,
		CurrentValue: user.CurrentValue,
		Status:       user.Status,
		Description:  user.Description,
		Records:      RecordsConvertToListResponse(user.Records),
	}
	return userRes
}

func GetInvestments(params *httpmodels.Filter) []models.Investments {
	users := []models.Investments{}
	query := db.DB.Preload("Companies.Currencies").Preload("Users").Preload("Instruments").Preload("Currencies").Preload("Records")
	query = utils.Filter(query, params)
	query.Find(&users)
	return users
}

func GetInvestmentsDetail(params *httpmodels.Filter) []models.Investments {
	users := []models.Investments{}
	query := db.DB.Preload("Records").Preload("Instruments")
	query = utils.Filter(query, params)
	query.Find(&users)
	return users
}

func GetInvestmentByFilters(params *httpmodels.Filter) *models.Investments {
	users := &models.Investments{}
	query := db.DB
	query = utils.Filter(query, params)
	query.First(&users)
	return users
}

func GetInvestmentByID(id interface{}) *models.Investments {

	user := models.Investments{}
	query := db.DB.Preload("Records").Where("id = ?", id)
	if query.First(&user).Error != nil {
		return nil
	}
	return &user
}

func CreateInvestment(params *httpmodels.CreateInvestmentRequest, user models.Users) (*models.Investments, error) {
	field := "companies_id"
	value := strconv.Itoa(int(params.Company))
	companiesShare := GetCompaniesShareByFilters(utils.CreateFilter(field, value))
	if companiesShare == nil {
		return nil, errors.New("No se encontro el valor de las acciones")
	}
	totalValue := (companiesShare.SharesValue * params.Shares)
	u := &models.Investments{
		UsersID:       user.ID,
		CompaniesID:   params.Company,
		InstrumentsID: params.Instrument,
		Shares:        params.Shares,
		InitValue:     totalValue,
		CurrentValue:  totalValue,
		Status:        true,
		Description:   params.Description,
		CurrenciesID:  companiesShare.Companies.Currencies.ID,
	}
	err := u.Create()
	if err == nil {
		_, err = CreateRecord(params, u)
		if err != nil {
			delete := utils.BindDelete(strconv.Itoa(int(u.ID)))
			DeleteInvestment(delete)
		}
	}
	return u, err
}

func UpdateInvestment(params *httpmodels.UpdateInvestmentRequest) (*models.Investments, error) {
	u := &models.Investments{}
	if db.DB.Where("id = ?", params.ID).First(u).Error != nil {
		return u, nil
	}
	field := "companies_id"
	value := strconv.Itoa(int(u.CompaniesID))
	companiesShare := GetCompaniesShareByFilters(utils.CreateFilter(field, value))
	if companiesShare == nil {
		return u, nil
	}
	currentValue := u.CurrentValue
	var totalValue float64
	if params.Shares != 0 {
		totalValue = (companiesShare.SharesValue * params.Shares) + currentValue
		params.Yield = 0
	} else {
		totalValue = (currentValue * (params.Yield / 100)) + currentValue
	}
	u.InstrumentsID = params.Instrument
	u.Shares = u.Shares + params.Shares
	u.CurrentValue = totalValue
	u.Status = params.Status
	u.Description = params.Description
	err := u.Save()
	if err == nil {
		UpdateInvestmentRecord(params, u, currentValue, totalValue)
	}
	return u, err
}

func DeleteInvestment(params *httpmodels.DeleteRequest) error {
	c := models.Investments{}
	if db.DB.Where("id = ?", params.ID).First(&c).Error != nil {
		return nil
	}
	err := DeleteRecord(params)
	if err == nil {
		err = c.Delete()
	}
	return err
}
