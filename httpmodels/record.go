package httpmodels

import (
	"time"
)

type CreateRecordRequest struct {
	Investment   uint    `json:"investment_id"`
	CurrentValue float64 `json:"current_value"`
	Yield        float64 `json:"yield"`
	TotalValue   bool    `json:"total_value"`
}

type CreateRecordResponse struct {
	Record *Record `json:"-"`
}

type GetRecordResponse struct {
	Record *Record `json:"-"`
}

type Record struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Investment   uint      `json:"investment_id"`
	CurrentValue float64   `json:"current_value"`
	Yield        float64   `json:"yield"`
	TotalValue   float64   `json:"total_value"`
	TypeRecords  int       `json:"type_records"`
}

type GetRecordDetail struct {
	MinDate      time.Time `json:"min_date"`
	MaxDate      time.Time `json:"max_date"`
	InvestmentID uint      `json:"investment_id"`
}

type GetRecordDetailResponse struct {
	Investment uint    `json:"investment_id"`
	InitValue  float64 `json:"init_value"`
	TotalValue float64 `json:"total_value"`
	Gain       float64 `json:"gain"`
}
