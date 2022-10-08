package models

import (
	"time"
)

type Transaction struct {
	ID            uint64    `json:"id"`
	UserID        uint      `json:"user_id"`
	User          User      `json:"user"`
	Address       string    `json:"address"`
	Payment       string    `json:"payment"`
	TotalPrice    float64   `json:"total_price"`
	ShippingPrice float64   `json:"shipping_price"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

type TransactionItems struct {
	ID            string      `json:"id"`
	UserID        uint        `json:"user_id"`
	User          User        `json:"user"`
	ProductID     uint        `json:"product_id"`
	Product       Product     `json:"product"`
	TransactionID string      `json:"transaction_id"`
	Transaction   Transaction `json:"transaction"`
	Quantity      int32       `json:"quantity"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
