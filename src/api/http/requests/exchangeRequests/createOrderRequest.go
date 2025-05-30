package exchangeRequests

type CreateOrderRequest struct {
	ExchangeID    uint    `json:"exchange_name"`
	ClientOrderID string  `json:"client_order_id"`
	Symbol        string  `json:"symbol"`
	Side          string  `json:"side"`
	Type          string  `json:"type"`
	Quantity      float64 `json:"quantity"`
	Price         float64 `json:"price"`
	TimeInForce   string  `json:"time_in_force"`
}
