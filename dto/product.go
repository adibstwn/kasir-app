package dto

type CreateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	CategoryId string  `json:"category_id"`
}

type UpdateProduct struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
