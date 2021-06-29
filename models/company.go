package models

type Companies struct {
	BaseModel
	Name         string     `json:"name" gorm:"column:name"`
	Currencies   Currencies `json:"currencies" belongs_to:"currencies"`
	CurrenciesID uint       `json:"currencies_id" gorm:"column:currencies_id"`
}
