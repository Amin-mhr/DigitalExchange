//go:build wireinject
// +build wireinject

package providers

import (
	"DigitalExchange/src/api/http/controllers"
	"DigitalExchange/src/database"
	"github.com/google/wire"
)

func InitializeServiceContainer() *ServiceContainer {
	wire.Build(
		database.GetInstance,
		ExchangeContainerSet,
		wire.Struct(new(ServiceContainer), "*"),
	)
	return nil
}

type ServiceContainer struct {
	ExchangeController *controllers.ExchangeController
}
