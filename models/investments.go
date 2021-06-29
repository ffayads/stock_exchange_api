package models

import "github.com/api/stock_exchange_api/config/db"

type Investments struct {
	BaseModel
	Users         Users        `json:"users" belongs_to:"users"`
	UsersID       uint         `json:"users_id" gorm:"column:users_id"`
	Companies     Companies    `json:"companies" belongs_to:"companies"`
	CompaniesID   uint         `json:"companies_id" gorm:"column:companies_id"`
	Instruments   Instruments  `json:"instruments" belongs_to:"instruments"`
	InstrumentsID uint         `json:"instruments_id" gorm:"column:instruments_id"`
	Currencies    Currencies   `json:"currencies" belongs_to:"currencies"`
	CurrenciesID  uint         `json:"currencies_id" gorm:"column:currencies_id"`
	Shares        float64      `json:"shares" gorm:"column:shares"`
	InitValue     float64      `json:"init_value" gorm:"column:init_value"`
	CurrentValue  float64      `json:"current_value" gorm:"column:current_value"`
	Status        bool         `json:"status" gorm:"column:status"`
	Description   string       `json:"description" gorm:"column:description;type:text"`
	Records       RecordsArray `has_many:"records" json:"-"`
}

type InvestmentsArray []Investments

func (model *Investments) Create() error {
	return db.DB.Create(model).Error
}

func (model *Investments) Save() error {
	return db.DB.Save(model).Error
}

func (model *Investments) Delete() error {
	return db.DB.Delete(model).Error
}
