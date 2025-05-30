package models

import "time"

type ExchangeModel struct {
	ID        uint      `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	APIKey    string    `json:"api_key" db:"api_key"`
	APISecret string    `json:"api_secret" db:"api_secret"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (*ExchangeModel) TableName() string {
	return "exchanges"
}
