package dto

type CreateCategoryReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryReq struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
