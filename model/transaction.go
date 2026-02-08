package model

import "time"

type Transaction struct {
	Id          string    `json:"id"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}
