package model

type Product struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      float64 `json:"stock"`
	IsDeleted  string  `json:"is_deleted"`
	IdCategory *string `json:"category_id,omitempty"`
	//UserCreator
	//CategoryName *string `json:"category_name"`
}
