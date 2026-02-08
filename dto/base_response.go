package dto

type BaseResponse struct {
	StatusCode    int         `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	ErrorMessage  *string     `json:"error_message,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	Pagination    *Pagination `json:"paging,omitempty"`
}

type Pagination struct {
	Page      int `json:"page"`
	MaxRow    int `json:"max_row"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
}
