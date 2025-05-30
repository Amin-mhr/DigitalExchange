package seeders

import (
	"DigitalExchange/src/database"
	"DigitalExchange/src/models"
	"log"
	"time"
)

func SeedBalances() {
	var db = database.GetInstance()

	balances := []*models.BalanceModel{
		{
			ID:         1,
			ExchangeID: 1,
			Asset:      "BTC",
			Free:       0.8,
			Locked:     0.2,
			UpdatedAt:  time.Date(2025, 5, 30, 18, 31, 0, 0, time.FixedZone("CEST", 2*60*60)),
		},
		{
			ID:         2,
			ExchangeID: 1,
			Asset:      "USDT",
			Free:       1500.0,
			Locked:     300.0,
			UpdatedAt:  time.Date(2025, 5, 30, 18, 31, 30, 0, time.FixedZone("CEST", 2*60*60)),
		},
		{
			ID:         3,
			ExchangeID: 2,
			Asset:      "ETH",
			Free:       2.5,
			Locked:     0.5,
			UpdatedAt:  time.Date(2025, 5, 30, 18, 32, 0, 0, time.FixedZone("CEST", 2*60*60)),
		},
		{
			ID:         4,
			ExchangeID: 2,
			Asset:      "XRP",
			Free:       2000.0,
			Locked:     500.0,
			UpdatedAt:  time.Date(2025, 5, 30, 18, 32, 30, 0, time.FixedZone("CEST", 2*60*60)),
		},
	}

	for _, balance := range balances {
		db.GetClient().FirstOrCreate(&balance, models.ExchangeModel{ID: balance.ID})
	}

	log.Println("Balance Seeder executed successfully.")
}
