package seeders

import (
	"DigitalExchange/src/database"
	"DigitalExchange/src/models"
	"log"
)

func SeedExchanges() {
	var db = database.GetInstance()

	exchanges := []*models.ExchangeModel{
		{
			ID:        1,
			Name:      "Binance",
			APIKey:    "APIKEY",
			APISecret: "APISECRET",
		},
		{
			ID:        2,
			Name:      "KuCoin",
			APIKey:    "APIKEY",
			APISecret: "APISECRET",
		},
	}

	for _, exchange := range exchanges {
		db.GetClient().FirstOrCreate(&exchange, models.ExchangeModel{ID: exchange.ID})
	}

	log.Println("Exchange Seeder executed successfully.")
}
