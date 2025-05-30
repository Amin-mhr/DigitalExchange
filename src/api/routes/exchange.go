package routes

import (
	wire "DigitalExchange/src/providers"
	"github.com/gin-gonic/gin"
)

func RegisterExchangeRoutes(router *gin.RouterGroup) {
	serviceContainer := wire.InitializeServiceContainer()

	matchRoutes := router.Group("exchange")
	{
		matchRoutes.POST("order", serviceContainer.ExchangeController.CreateOrder)

		matchRoutes.DELETE("order/:id", serviceContainer.ExchangeController.DeleteOrder)

		matchRoutes.GET("balance/:id", serviceContainer.ExchangeController.GetBalance)

		matchRoutes.GET("orderbook/:id", serviceContainer.ExchangeController.GetOrderBook)

	}
}
