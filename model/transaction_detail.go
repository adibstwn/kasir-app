package model

type TransactionDetail struct {
	Id            string  `json:"id"`
	TransactionId string  `json:"transaction_id"`
	ProductId     string  `json:"product_id"`
	Quantity      float64 `json:"quantity"`
	Amount        float64 `json:"amount"`
}

type TransactionReportDb struct {
	Id            string  `json:"id"`
	TransactionId string  `json:"transaction_id"`
	ProductId     string  `json:"product_id"`
	Quantity      float64 `json:"quantity"`
	Amount        float64 `json:"amount"`
	ProductName   string  `json:"product_name"`
}
