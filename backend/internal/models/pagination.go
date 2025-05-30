// internal/models/pagination.go
package models

// PaginationParams represents the parameters for pagination
type PaginationParams struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

// PaginatedResult represents a paginated result set
type PaginatedResult struct {
	Items      interface{} `json:"items"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalPages int         `json:"totalPages"`
}

// NewPaginatedResult creates a new paginated result
func NewPaginatedResult(items interface{}, total int64, params PaginationParams) PaginatedResult {
	totalPages := int(total) / params.Limit
	if int(total)%params.Limit > 0 {
		totalPages++
	}

	return PaginatedResult{
		Items:      items,
		Total:      total,
		Page:       params.Page,
		Limit:      params.Limit,
		TotalPages: totalPages,
	}
}