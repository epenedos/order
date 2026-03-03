package pagination

import (
	"net/http"
	"strconv"
)

const (
	DefaultPage  = 1
	DefaultLimit = 20
	MaxLimit     = 100
)

type Params struct {
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"-"`
}

func FromRequest(r *http.Request) Params {
	page := parseIntParam(r, "page", DefaultPage)
	limit := parseIntParam(r, "limit", DefaultLimit)

	if page < 1 {
		page = DefaultPage
	}
	if limit < 1 {
		limit = DefaultLimit
	}
	if limit > MaxLimit {
		limit = MaxLimit
	}

	return Params{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}
}

type PagedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalCount int         `json:"total_count"`
	TotalPages int         `json:"total_pages"`
}

func NewPagedResponse(data interface{}, params Params, totalCount int) PagedResponse {
	totalPages := totalCount / params.Limit
	if totalCount%params.Limit > 0 {
		totalPages++
	}
	return PagedResponse{
		Data:       data,
		Page:       params.Page,
		Limit:      params.Limit,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}
}

func parseIntParam(r *http.Request, key string, defaultVal int) int {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultVal
	}
	parsed, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return parsed
}
