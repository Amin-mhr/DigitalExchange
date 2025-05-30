package repositories

import (
	"DigitalExchange/src/database"
	"DigitalExchange/src/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type IExchangeRepository interface {
	CreateOrder(Order *models.OrderModel) (*models.OrderModel, error)
	GetOrderById(OrderID uint) (*models.OrderModel, error)
	DeleteOrder(Order *models.OrderModel) error
	GetBalanceListExchange(exchangeID uint) ([]*models.BalanceModel, error)
	GetOrderBookListExchange(exchangeID uint) ([]*models.OrderBook, error)
}

type ExchangeRepository struct {
	IDatabaseHandler *database.Database
}

func (repository *ExchangeRepository) CreateOrder(Order *models.OrderModel) (*models.OrderModel, error) {
	result := repository.IDatabaseHandler.GetClient().Create(&Order)
	if result.Error != nil {
		return nil, fmt.Errorf("Order creation failed: %s", result.Error.Error())
	}
	return Order, nil
}

func (repository *ExchangeRepository) GetOrderById(OrderID uint) (*models.OrderModel, error) {
	var Order models.OrderModel
	result := repository.IDatabaseHandler.GetClient().
		First(&Order, "id = ?", OrderID)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("Order get by id failed: %s", result.Error.Error())
	}

	return &Order, nil
}

func (repository *ExchangeRepository) DeleteOrder(Order *models.OrderModel) error {
	result := repository.IDatabaseHandler.GetClient().Delete(&Order)
	if result.Error != nil {
		return fmt.Errorf("Order delete failed: %s", result.Error.Error())
	}
	return nil
}

func (repository *ExchangeRepository) GetBalanceListExchange(exchangeID uint) ([]*models.BalanceModel, error) {
	var balances []*models.BalanceModel
	result := repository.IDatabaseHandler.GetClient().Where("exchange_id = ?", exchangeID).Find(&balances)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch balances: %s", result.Error.Error())
	}
	return balances, nil
}

func (repository *ExchangeRepository) GetOrderBookListExchange(exchangeID uint) ([]*models.OrderBook, error) {
	var orderBooks []*models.OrderBook
	result := repository.IDatabaseHandler.GetClient().Where("exchange_id = ?", exchangeID).Find(&orderBooks)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch orderBooks: %s", result.Error.Error())
	}
	return orderBooks, nil
}
