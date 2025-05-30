package seeders

import (
	"DigitalExchange/src/database"
	"DigitalExchange/src/models"
	"log"
	"time"
)

func SeedOrders() {
	var db = database.GetInstance()

	orders := []*models.OrderModel{
		{
			ID:              1,
			ExchangeID:      1,
			ClientOrderID:   "buy-btc-limit-20250530-001",
			Symbol:          "BTC/USDT",
			Side:            "buy",
			Type:            "limit",
			Quantity:        0.2,
			Price:           60000.0,
			TimeInForce:     "GTC",
			Status:          "pending",
			ExchangeOrderID: "100001",
			CreatedAt:       time.Date(2025, 5, 30, 18, 25, 0, 0, time.FixedZone("CEST", 2*60*60)),
			UpdatedAt:       time.Date(2025, 5, 30, 18, 25, 0, 0, time.FixedZone("CEST", 2*60*60)),
		},
		{
			ID:              2,
			ExchangeID:      1,
			ClientOrderID:   "sell-eth-market-20250530-002",
			Symbol:          "ETH/USDT",
			Side:            "sell",
			Type:            "market",
			Quantity:        1.5,
			Price:           0.0,
			TimeInForce:     "GTC",
			Status:          "pending",
			ExchangeOrderID: "100002",
			CreatedAt:       time.Date(2025, 5, 30, 18, 25, 30, 0, time.FixedZone("CEST", 2*60*60)),
			UpdatedAt:       time.Date(2025, 5, 30, 18, 25, 30, 0, time.FixedZone("CEST", 2*60*60)),
		},
		{
			ID:              3,
			ExchangeID:      2,
			ClientOrderID:   "buy-xrp-limit-20250530-003",
			Symbol:          "LTC/BTC",
			Side:            "buy",
			Type:            "limit",
			Quantity:        1000.0,
			Price:           0.5,
			TimeInForce:     "GTC",
			Status:          "pending",
			ExchangeOrderID: "200001",
			CreatedAt:       time.Date(2025, 5, 30, 18, 26, 0, 0, time.FixedZone("CEST", 2*60*60)),
			UpdatedAt:       time.Date(2025, 5, 30, 18, 26, 0, 0, time.FixedZone("CEST", 2*60*60)),
		},
		{
			ID:              4,
			ExchangeID:      2,
			ClientOrderID:   "sell-ltc-limit-20250530-004",
			Symbol:          "LTC/BTC",
			Side:            "sell",
			Type:            "limit",
			Quantity:        10.0,
			Price:           0.0025,
			TimeInForce:     "GTC",
			Status:          "pending",
			ExchangeOrderID: "200002",
			CreatedAt:       time.Date(2025, 5, 30, 18, 26, 30, 0, time.FixedZone("CEST", 2*60*60)),
			UpdatedAt:       time.Date(2025, 5, 30, 18, 26, 30, 0, time.FixedZone("CEST", 2*60*60)),
		},
	}

	for _, order := range orders {
		db.GetClient().FirstOrCreate(&order, models.ExchangeModel{ID: order.ID})
	}

	log.Println("order Seeder executed successfully.")
}
