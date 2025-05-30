package services

import (
	"DigitalExchange/src/api/http/requests/exchangeRequests"
	"DigitalExchange/src/models"
	"DigitalExchange/src/repositories"
)

type IExchangeService interface {
	CreateOrder(request *exchangeRequests.CreateOrderRequest) (*models.OrderModel, error)
	DeleteOrder(OrderID uint) (*models.OrderModel, error)
	GetOrderById(OrderID uint) (*models.OrderModel, error)
	GetBalance(exchangeID uint) ([]*models.BalanceModel, error)
	GetOrderBook(exchangeID uint) ([]*models.OrderBook, error)
}

type ExchangeService struct {
	IExchangeRepository repositories.IExchangeRepository
}

func (service *ExchangeService) CreateOrder(request *exchangeRequests.CreateOrderRequest) (*models.OrderModel, error) {

	Order := &models.OrderModel{
		ExchangeID:    request.ExchangeID,
		ClientOrderID: request.ClientOrderID,
		Symbol:        request.Symbol,
		Side:          request.Side,
		Type:          request.Type,
		Quantity:      request.Quantity,
		Price:         request.Price,
		TimeInForce:   request.TimeInForce,
		Status:        "pending",
	}

	OrderOrm, err := service.IExchangeRepository.CreateOrder(Order)
	if err != nil {
		return nil, err
	}

	return OrderOrm, nil
}

func (service *ExchangeService) GetOrderById(OrderID uint) (*models.OrderModel, error) {
	return service.IExchangeRepository.GetOrderById(OrderID)
}

func (service *ExchangeService) DeleteOrder(orderId uint) error {
	Order, err := service.GetOrderById(orderId)

	if err != nil {
		return err
	}

	return service.IExchangeRepository.DeleteOrder(Order)
}

func (service *ExchangeService) GetBalance(exchangeID uint) ([]*models.BalanceModel, error) {
	balances, err := service.IExchangeRepository.GetBalanceListExchange(exchangeID)
	if err != nil {
		return nil, err
	}
	return balances, nil
}

func (service *ExchangeService) GetOrderBook(exchangeID uint) ([]*models.OrderBook, error) {
	orderBooks, err := service.IExchangeRepository.GetOrderBookListExchange(exchangeID)
	if err != nil {
		return nil, err
	}
	return orderBooks, nil
}
