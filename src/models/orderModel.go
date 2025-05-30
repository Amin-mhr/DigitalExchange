package models

import "time"

type OrderModel struct {
	ID              uint      `json:"id" db:"id"`
	ExchangeID      uint      `json:"exchange_id" db:"exchange_id" gorm:"foreignKey:ID;references:ID"`
	ClientOrderID   string    `json:"client_order_id" db:"client_order_id"`
	Symbol          string    `json:"symbol" db:"symbol"`
	Side            string    `json:"side" db:"side"`
	Type            string    `json:"type" db:"type"`
	Quantity        float64   `json:"quantity" db:"quantity"`
	Price           float64   `json:"price" db:"price"`
	TimeInForce     string    `json:"time_in_force" db:"time_in_force"`
	Status          string    `json:"status" db:"status"`
	ExchangeOrderID string    `json:"exchange_order_id" db:"exchange_order_id"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

func (*OrderModel) TableName() string {
	return "orders"
}
