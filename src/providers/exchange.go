package providers

import (
	"DigitalExchange/src/api/http/controllers"
	"DigitalExchange/src/database"
	"DigitalExchange/src/repositories"
	"DigitalExchange/src/services"
	"github.com/google/wire"
)

// ProvideMatchRepository is a Wire provider function that constructs a MatchRepository.
func ProvideExchangeRepository(db *database.Database) *repositories.ExchangeRepository {
	return &repositories.ExchangeRepository{
		IDatabaseHandler: db,
	}
}

// ProvideMatchService is a Wire provider function that constructs a MatchService.
func ProvideExchangeService(exchangeRepository *repositories.ExchangeRepository) *services.ExchangeService {
	return &services.ExchangeService{
		IExchangeRepository: exchangeRepository,
	}
}

// ProvideMatchController is a Wire provider function that constructs an MatchController.
func ProvideExchangeController(ExchangeService *services.ExchangeService) *controllers.ExchangeController {
	return &controllers.ExchangeController{
		IExchangeService: ExchangeService,
	}
}

// MatchContainerSet is Controllers, services, repositories provider set
var ExchangeContainerSet = wire.NewSet(
	ProvideExchangeController,
	ProvideExchangeService,
	ProvideExchangeRepository,
)
