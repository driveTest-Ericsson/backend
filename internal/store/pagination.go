package store

import (
	"net/http"
	"strconv"
	"time"
)

type PaginatedCellQuery struct {
	Limit  int    `json:"limit" validate:"gte=1,lte=20"`
	Offset int    `json:"offset" validate:"gte=0"`
	Sort   string `json:"sort" validate:"oneof=asc desc"`
	Search string `json:"search" validate:"max=100"`
	Since  string `json:"since"`
	Until  string `json:"until"`
}

func (cq *PaginatedCellQuery) Parse(r *http.Request) (*PaginatedCellQuery, error) {

	qs := r.URL.Query()

	limit := qs.Get("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return cq, nil
		}

		cq.Limit = l
	}

	offset := qs.Get("offset")
	if offset != "" {
		o, err := strconv.Atoi(offset)
		if err != nil {
			return cq, nil
		}

		cq.Offset = o
	}

	sort := qs.Get("sort")
	if sort != "" {
		cq.Sort = sort
	}

	search := qs.Get("search")
	if search != "" {
		cq.Search = search
	}

	since := qs.Get("since")
	if since != "" {
		cq.Since = parseTime(since)
	}

	until := qs.Get("until")
	if until != "" {
		cq.Until = parseTime(until)
	}

	return cq, nil
}

func parseTime(s string) string {
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
		return ""
	}

	return t.Format(time.DateTime)
}
