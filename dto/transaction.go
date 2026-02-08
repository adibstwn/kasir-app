package dto

type CheckoutItem struct {
	ProductId string  `json:"product_id"`
	Quantity  float64 `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

type TransactionReport struct {
	TotalRevenue      float64           `json:"total_revenue"`
	TotalTransaction  int64             `json:"total_transaction"`
	BestSellerProduct BestSellerProduct `json:"produk_terlaris"`
}
type BestSellerProduct struct {
	Name    string  `json:"name"`
	QtySell float64 `json:"qty_sell"`
}
