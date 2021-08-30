package utility

import (
	"strconv"
)

type ModelPagination struct {
	Page  int64  `json:"p" query:"p"`
	Limit int64  `json:"l" query:"l"`
	Query string `json:"q" query:"q"`
	Sort  string `json:"s" query:"s"`
}

func ConvertPagination(page, limit, query []byte) ModelPagination {
	valPage, _ := strconv.Atoi(string(page))
	valLimit, _ := strconv.Atoi(string(limit))

	return ModelPagination{
		Limit: int64(valLimit),
		Page:  int64(valPage),
		Query: string(query),
	}
}

func (a ModelPagination) ValidatePagination() (output ModelPagination) {
	output.Page = a.Page
	output.Limit = a.Limit
	output.Query = a.Query
	if output.Limit == 0 {
		output.Limit = 10
	}

	if output.Page <= 1 {
		output.Page = 0
	} else {
		output.Page = (output.Page - 1) * output.Limit
	}

	if a.Sort == "" {
		output.Sort = "desc"
	} else {
		output.Sort = a.Sort
	}
	return
}
