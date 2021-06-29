package utils

import (
	httpmodels "github.com/api/stock_exchange_api/httpmodels"
	"gorm.io/gorm"
)

func CreateFilter(field, value string) *httpmodels.Filter {
	if value == "" || field == "" {
		return &httpmodels.Filter{}
	}
	conditions := []httpmodels.Conditions{}
	values := []string{value}
	condition := &httpmodels.Conditions{
		Field: field,
		Value: values,
	}
	conditions = append(conditions, *condition)
	filter := &httpmodels.Filter{
		Conditions: conditions,
	}
	return filter
}

func Filter(query *gorm.DB, filter *httpmodels.Filter) *gorm.DB {
	for _, fields := range filter.Conditions {
		if len(fields.Value) > 0 && fields.Field != "" {
			var count int
			for _, value := range fields.Value {
				if count == 0 {
					if fields.Operation == "equal" {
						query = query.Where(fields.Field+" = ?", value)
					} else if fields.Operation == "like" {
						query = query.Where(fields.Field+" like ?", "%"+value+"%")
					} else {
						query = query.Where(fields.Field+" like ?", "%"+value+"%")
					}
				} else {
					if fields.Operation == "equal" {
						query = query.Or(fields.Field+" = ?", value)
					} else if fields.Operation == "like" {
						query = query.Or(fields.Field+" like ?", "%"+value+"%")
					} else {
						query = query.Or(fields.Field+" like ?", "%"+value+"%")
					}
				}
				count++
			}
		}
	}
	return query
}

func OrderQuery(query *gorm.DB, params *httpmodels.Filter, orderQuery string, limit, offset, page int) (*gorm.DB, int, int, int) {
	if params != nil && params.Order != nil {
		orderQuery = *params.Order
		if params.OrderType != nil {
			orderQuery = orderQuery + " " + *params.OrderType
		}
	}
	query = query.Order(orderQuery)
	query, limit, offset, page = LimitQuery(query, params, limit, offset, page)
	return query, limit, offset, page
}

func LimitQuery(query *gorm.DB, params *httpmodels.Filter, limit, offset, page int) (*gorm.DB, int, int, int) {
	if params != nil && params.Limit != nil && *params.Limit > 0 {
		limit = *params.Limit
	}
	query = query.Limit(limit)
	query, offset, page = OffsetQuery(query, params, limit, offset, page)
	return query, limit, offset, page
}

func OffsetQuery(query *gorm.DB, params *httpmodels.Filter, limit, offset, page int) (*gorm.DB, int, int) {
	if params != nil && params.Page != nil && *params.Page > 1 {
		page = *params.Page
	}
	offset = (page - 1) * limit
	query = query.Offset(offset)
	return query, offset, page
}
