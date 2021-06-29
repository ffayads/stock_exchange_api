package repository

import (
	"github.com/api/stock_exchange_api/config/db"
	"github.com/api/stock_exchange_api/httpmodels"
	"github.com/api/stock_exchange_api/models"
	"github.com/api/stock_exchange_api/utils"
)

func GetCompaniesShareByFilters(params *httpmodels.Filter) *models.CompaniesShare {
	model := &models.CompaniesShare{}
	query := db.DB.Preload("Companies.Currencies").Where("status = ?", true)
	query = utils.Filter(query, params)
	query.First(&model)
	return model
}
