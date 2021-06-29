package httpmodels

import (
	"time"
)

type CreateInvestmentRequest struct {
	Company     uint    `json:"company_id"`
	Instrument  uint    `json:"instrument_id"`
	Shares      float64 `json:"shares"`
	Status      bool    `json:"status"`
	Description string  `json:"description"`
}

type CreateInvestmentResponse struct {
	Investment *Investment `json:"-"`
}

type UpdateInvestmentRequest struct {
	ID          uint    `json:"id"`
	Instrument  uint    `json:"instrument_id"`
	Shares      float64 `json:"shares"`
	Status      bool    `json:"status"`
	Description string  `json:"description"`
	Yield       float64 `json:"yield"`
}

type UpdateInvestmentResponse struct {
	Investment *Investment `json:"-"`
}

type GetInvestmentResponse struct {
	Investment *Investment `json:"-"`
}

type Investment struct {
	ID           uint        `json:"id"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	User         *User       `json:"user"`
	Company      *Company    `json:"company"`
	Instrument   *Instrument `json:"instrument"`
	Currency     *Currency   `json:"currency"`
	Shares       float64     `json:"shares"`
	InitValue    float64     `json:"init_value"`
	CurrentValue float64     `json:"current_value"`
	Status       bool        `json:"status"`
	Description  string      `json:"description"`
	Records      []*Record   `json:"records"`
}
