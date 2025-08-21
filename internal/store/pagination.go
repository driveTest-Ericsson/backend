package store

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

type OrderByField struct {
	Field     string `json:"field" validate:"oneof=cell_tech frequency_band rxlev rscp rsrp sinr generated_at"`
	Direction string `json:"direction" validate:"oneof=asc desc"`
}

type PaginatedCellQuery struct {
	Limit          int      `json:"limit" validate:"gte=1,lte=300"`
	Offset         int      `json:"offset" validate:"gte=0"`
	Search         string   `json:"search" validate:"max=100"`
	Since          string   `json:"since"`
	Until          string   `json:"until"`
	OrderBy        string   `json:"order_by"`
	CellTech       []string `json:"cell_tech"`
	FrequencyBands []string `json:"frequency_band"`
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

	// sort := qs.Get("sort")
	// if sort != "" {
	// 	cq.Sort = sort
	// }

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

	cellTech := qs.Get("cell_tech")
	if cellTech != "" {
		cq.CellTech = splitCommaSeparated(cellTech)
	}

	frequencyBand := qs.Get("frequency_band")
	if frequencyBand != "" {
		cq.FrequencyBands = splitCommaSeparated(frequencyBand)
	}

	orderBy := qs.Get("order_by")
	orderByString := "generated_at desc"
	if orderBy != "" {
		parts := splitCommaSeparated(orderBy)

		for _, part := range parts {
			subParts := strings.Split(part, ":")
			if len(subParts) != 2 {
				continue
			}

			field := strings.TrimSpace(subParts[0])
			direction := strings.TrimSpace(subParts[1])

			if field == "" || direction == "" {
				continue
			}

			orderByString += ", " + field + " " + direction
		}

		cq.OrderBy = orderByString
	}

	return cq, nil
}

func splitCommaSeparated(s string) []string {
	parts := strings.Split(s, ",")
	var result []string
	for _, part := range parts {
		p := strings.TrimSpace(part)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}

func parseTime(s string) string {
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
		return ""
	}

	return t.Format(time.DateTime)
}
