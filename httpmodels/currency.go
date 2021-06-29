package httpmodels

import (
	"time"
)

type Currency struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Acronym   string    `json:"acronym"`
}
