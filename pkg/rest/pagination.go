package rest

import (
	"context"
)

// Page type.
type Page struct {
	List       interface{} `json:"page_list"`
	Page       int64       `json:"page_page"`
	Limit      int64       `json:"page_limit"`
	Pages      int64       `json:"page_pages"`
	SizeOfList int64       `json:"page_size_of_list"`
	Catalogues interface{} `json:"catalogues,omitempty"`
	Filters    interface{} `json:"filters,omitempty"`
}

// NewPagev2 convenience.
func NewPagev2(ctx context.Context, list interface{}, size int64, limit, page int64) *Page {
	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}
	total := size / limit
	rest := size % limit
	if rest > 0 {
		total++
	}
	p := &Page{
		List:       list,
		Page:       page,
		Limit:      limit,
		Pages:      total,
		SizeOfList: size,
	}
	return p
}

// NewPageWithinCataloguesAndFilters convenience.
func NewPageWithinCataloguesAndFilters(ctx context.Context, list, catalogues interface{}, filters interface{}, size int64, limit, page int64) *Page {

	total := size / limit
	rest := size % limit
	if rest > 0 {
		total++
	}
	p := &Page{
		List:       list,
		Page:       page,
		Limit:      limit,
		Pages:      total,
		SizeOfList: size,
		Catalogues: catalogues,
		Filters:    filters,
	}
	return p
}
