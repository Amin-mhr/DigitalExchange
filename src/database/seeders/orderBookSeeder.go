package seeders

import (
	"DigitalExchange/src/database"
	"DigitalExchange/src/models"
	"log"
	"time"
)

func SeedOrderBooks() {
	var db = database.GetInstance()

	orderBooks := []*models.OrderBook{
		{
			ID:          1,
			ExchangeID:  1,
			Symbol:      "BTC/USDT",
			BidPrice:    59950.0,
			BidQuantity: 0.3,
			AskPrice:    60050.0,
			AskQuantity: 0.2,
			Timestamp:   time.Date(2025, 5, 30, 18, 28, 0, 0, time.FixedZone("CEST", 2*60*60)),
			UpdatedAt:   time.Date(2025, 5, 30, 18, 28, 0, 0, time.FixedZone("CEST", 2*60*60)),
		},
		{
			ID:          2,
			ExchangeID:  1,
			Symbol:      "ETH/USDT",
			BidPrice:    3500.0,
			BidQuantity: 1.2,
			AskPrice:    3510.0,
			AskQuantity: 0.8,
			Timestamp:   time.Date(2025, 5, 30, 18, 28, 30, 0, time.FixedZone("CEST", 2*60*60)),
			UpdatedAt:   time.Date(2025, 5, 30, 18, 28, 30, 0, time.FixedZone("CEST", 2*60*60)),
		},
		{
			ID:          3,
			ExchangeID:  2,
			Symbol:      "LTC/BTC",
			BidPrice:    0.49,
			BidQuantity: 5000.0,
			AskPrice:    0.51,
			AskQuantity: 3000.0,
			Timestamp:   time.Date(2025, 5, 30, 18, 29, 0, 0, time.FixedZone("CEST", 2*60*60)),
			UpdatedAt:   time.Date(2025, 5, 30, 18, 29, 0, 0, time.FixedZone("CEST", 2*60*60)),
		},
		{
			ID:          4,
			ExchangeID:  2,
			Symbol:      "LTC/BTC",
			BidPrice:    0.0024,
			BidQuantity: 15.0,
			AskPrice:    0.0026,
			AskQuantity: 10.0,
			Timestamp:   time.Date(2025, 5, 30, 18, 29, 30, 0, time.FixedZone("CEST", 2*60*60)),
			UpdatedAt:   time.Date(2025, 5, 30, 18, 29, 30, 0, time.FixedZone("CEST", 2*60*60)),
		},
	}

	for _, orderBook := range orderBooks {
		db.GetClient().FirstOrCreate(&orderBook, models.ExchangeModel{ID: orderBook.ID})
	}

	log.Println("OrderBooks Seeder executed successfully.")
}
