package repository

import (
	"github.com/api/stock_exchange_api/config/db"
	"github.com/api/stock_exchange_api/httpmodels"
	"github.com/api/stock_exchange_api/models"
	"github.com/api/stock_exchange_api/utils"
)

func companiesConvertToListResponse(companies []models.Companies) []*httpmodels.Company {
	companiesResp := []*httpmodels.Company{}
	for _, company := range companies {
		companiesResp = append(companiesResp, CompanyConvertToResponse(&company))
	}
	return companiesResp
}

func CompanyConvertToResponse(company *models.Companies) *httpmodels.Company {

	companyRes := &httpmodels.Company{
		ID:           company.ID,
		CreatedAt:    company.CreatedAt,
		UpdatedAt:    company.UpdatedAt,
		CurrenciesID: company.CurrenciesID,
		Name:         company.Name,
	}
	return companyRes
}

func GetCompanyByFilters(params *httpmodels.Filter) *models.Companies {
	model := &models.Companies{}
	query := db.DB
	query = utils.Filter(query, params)
	query.First(&model)
	return model
}
