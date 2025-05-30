package models

import "time"

type BalanceModel struct {
	ID         uint      `json:"id" db:"id"`
	ExchangeID uint      `json:"exchange_id" db:"exchange_id" gorm:"foreignKey:ID;references:ID"`
	Asset      string    `json:"asset" db:"asset"`
	Free       float64   `json:"free" db:"free"`
	Locked     float64   `json:"locked" db:"locked"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

func (*BalanceModel) TableName() string {
	return "balances"
}
