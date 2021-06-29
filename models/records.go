package models

import "github.com/api/stock_exchange_api/config/db"

var RECORD_YIELD = 0
var RECORD_SHARE = 1

type Records struct {
	BaseModel
	Investments   Investments `json:"investments" belongs_to:"investments"`
	InvestmentsID uint        `json:"investments_id" gorm:"column:investments_id"`
	CurrentValue  float64     `json:"current_value" gorm:"column:current_value"`
	Yield         float64     `json:"yield" gorm:"column:yield"`
	TotalValue    float64     `json:"total_value" gorm:"column:total_value"`
	TypeRecords   int         `json:"type_record" gorm:"column:type_record"`
}

type RecordsArray []Records

func (model *Records) Create() error {
	return db.DB.Create(model).Error
}

func (model *Records) Save() error {
	return db.DB.Save(model).Error
}

func (model *Records) Delete() error {
	return db.DB.Delete(model).Error
}
