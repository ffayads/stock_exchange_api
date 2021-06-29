package httpmodels

import (
	"time"
)

type Company struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Currency     Currency  `json:"currency"`
	CurrenciesID uint      `json:"currencies_id"`
	Name         string    `json:"name"`
}
