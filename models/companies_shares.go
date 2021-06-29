package models

type CompaniesShare struct {
	BaseModel
	Companies   Companies `json:"companies" belongs_to:"companies"`
	CompaniesID uint      `json:"companies_id" gorm:"column:companies_id"`
	SharesValue float64   `json:"shares_values" gorm:"column:shares_values"`
	Status      bool      `json:"status" gorm:"column:status"`
}
