package repository

import (
	"github.com/api/stock_exchange_api/config/db"
	"github.com/api/stock_exchange_api/httpmodels"
	"github.com/api/stock_exchange_api/models"
	"github.com/api/stock_exchange_api/utils"
)

func InstrumentsConvertToListResponse(Instruments []models.Instruments) []*httpmodels.Instrument {
	instrumentsResp := []*httpmodels.Instrument{}
	for _, instrument := range Instruments {
		instrumentsResp = append(instrumentsResp, InstrumentConvertToResponse(&instrument))
	}
	return instrumentsResp
}

func InstrumentConvertToResponse(Instrument *models.Instruments) *httpmodels.Instrument {

	instrumentRes := &httpmodels.Instrument{
		ID:        Instrument.ID,
		CreatedAt: Instrument.CreatedAt,
		UpdatedAt: Instrument.UpdatedAt,
		Name:      Instrument.Name,
	}
	return instrumentRes
}

func GetInstrumentByFilters(params *httpmodels.Filter) *models.Instruments {
	model := &models.Instruments{}
	query := db.DB
	query = utils.Filter(query, params)
	query.First(&model)
	return model
}
