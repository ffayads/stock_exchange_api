package models

type Currencies struct {
	BaseModel
	Name    string `json:"name" gorm:"column:name"`
	Acronym string `json:"acronym" gorm:"column:acronym"`
}
