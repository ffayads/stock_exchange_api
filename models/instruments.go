package models

type Instruments struct {
	BaseModel
	Name string `json:"name" gorm:"column:name"`
}
