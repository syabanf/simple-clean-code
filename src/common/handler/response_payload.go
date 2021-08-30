package handler

type defaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func DefaultResponse(data interface{}, err error) (resp defaultResponse) {
	resp = defaultResponse{
		Message: "Success",
		Data:    data,
	}
	if err != nil {
		resp.Message = err.Error()
	}
	return
}

type responsePagination struct {
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	CurrentPage int         `json:"current_page"`
	Limit       int         `json:"limit"`
	TotalPage   int         `json:"total_page"`
	TotalData   int         `json:"total_data"`
}

func PaginationResponse(data interface{}, err error, page, limit, pageTotal, totalData int) (resp responsePagination) {
	resp = responsePagination{
		Message:     "Success",
		Data:        data,
		CurrentPage: page,
		Limit:       limit,
		TotalPage:   pageTotal,
		TotalData:   totalData,
	}
	if err != nil {
		resp.Message = err.Error()
	}
	return
}
