package controllers

import (
	"DigitalExchange/src/api/http/requests/exchangeRequests"
	"DigitalExchange/src/api/http/response"
	"DigitalExchange/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ExchangeController struct {
	IExchangeService *services.ExchangeService
}

func (controller *ExchangeController) DeleteOrder(c *gin.Context) {
	idStr := c.Param("id")

	uid, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Api(c).Send()
		return
	}
	id := uint(uid)

	err = controller.IExchangeService.DeleteOrder(id)

	if err != nil {
		response.Api(c).
			SetStatusCode(http.StatusNotFound).
			SetMessage(err.Error()).
			Send()
		return
	}

	response.Api(c).
		SetStatusCode(http.StatusOK).
		SetMessage("request-successful").
		Send()
	return
}

func (controller *ExchangeController) CreateOrder(c *gin.Context) {

	var req exchangeRequests.CreateOrderRequest

	// Bind check payload.
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Api(c).Send()
		return
	}

	//// Validate the payload.
	//if err := validator.Validate(&req, c.GetString("locale")); err != nil {
	//	response.Api(c).
	//		SetStatusCode(http.StatusUnprocessableEntity).
	//		SetErrors(err).
	//		Send()
	//	return
	//}

	// Pass the Match to service.
	Order, err := controller.IExchangeService.CreateOrder(&req)

	if err != nil {
		response.Api(c).
			SetMessage(err.Error()).
			Send()
		return
	}

	// Return response.
	response.Api(c).
		SetStatusCode(http.StatusCreated).
		SetData(map[string]interface{}{
			"Order": Order,
		}).
		SetMessage("request-successful").
		Send()
}

func (controller *ExchangeController) GetBalance(c *gin.Context) {
	idStr := c.Param("id")

	uid, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Api(c).Send()
		return
	}
	id := uint(uid)

	balanceList, err := controller.IExchangeService.GetBalance(id)

	if err != nil {
		response.Api(c).
			SetMessage(err.Error()).
			Send()
		return
	}

	// Return response.
	response.Api(c).
		SetStatusCode(http.StatusCreated).
		SetData(map[string]interface{}{
			"BalanceList": balanceList,
		}).
		SetMessage("request-successful").
		Send()
}

func (controller *ExchangeController) GetOrderBook(c *gin.Context) {
	idStr := c.Param("id")

	uid, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Api(c).Send()
		return
	}
	id := uint(uid)

	balanceList, err := controller.IExchangeService.GetOrderBook(id)

	if err != nil {
		response.Api(c).
			SetMessage(err.Error()).
			Send()
		return
	}

	// Return response.
	response.Api(c).
		SetStatusCode(http.StatusCreated).
		SetData(map[string]interface{}{
			"BalanceList": balanceList,
		}).
		SetMessage("request-successful").
		Send()
}
