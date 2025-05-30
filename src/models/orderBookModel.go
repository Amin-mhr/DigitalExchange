package models

import "time"

type OrderBook struct {
	ID          uint      `json:"id" db:"id"`
	ExchangeID  uint      `json:"exchange_id" db:"exchange_id" gorm:"foreignKey:ID;references:ID"`
	Symbol      string    `json:"symbol" db:"symbol"`
	BidPrice    float64   `json:"bid_price" db:"bid_price"`
	BidQuantity float64   `json:"bid_quantity" db:"bid_quantity"`
	AskPrice    float64   `json:"ask_price" db:"ask_price"`
	AskQuantity float64   `json:"ask_quantity" db:"ask_quantity"`
	Timestamp   time.Time `json:"timestamp" db:"timestamp"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (*OrderBook) TableName() string {
	return "order_books"
}
