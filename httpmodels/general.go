package httpmodels

type Filter struct {
	Order      *string      `json:"order"`
	OrderType  *string      `json:"order_type"`
	Limit      *int         `json:"limit"`
	Page       *int         `json:"page"`
	Conditions []Conditions `json:"conditions"`
	Value      *string      `json:"value"`
}

type Conditions struct {
	Field     string   `json:"fields"`
	Value     []string `json:"value"`
	Operation string   `json:"operation" default:"equal"`
}

type DeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}
