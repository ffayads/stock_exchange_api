package helpers

import (
	"errors"

	"github.com/api/stock_exchange_api/httpmodels"
	"github.com/api/stock_exchange_api/models"
	"github.com/api/stock_exchange_api/repository"
)

func CreateInvestment(i *httpmodels.CreateInvestmentRequest, user *models.Users) (*models.Investments, error) {
	investment, err := repository.CreateInvestment(i, *user)
	if err != nil {
		return nil, errors.New("No se pudo crear el usuario")
	}
	return investment, nil
}

func UpdateInvestment(i *httpmodels.UpdateInvestmentRequest) (*models.Investments, error) {
	investment, err := repository.UpdateInvestment(i)
	return investment, err
}

func DeleteInvestment(params *httpmodels.DeleteRequest) error {
	err := repository.DeleteInvestment(params)
	return err
}

func GetInvestments(values *httpmodels.Filter) []*httpmodels.Investment {
	return repository.InvestmentsConvertToListResponse(repository.GetInvestments(values))
}

func GetInvestmentsDetail(values *httpmodels.Filter) []*httpmodels.Investment {
	return repository.InvestmentsDetailConvertToListResponse(repository.GetInvestmentsDetail(values))
}
