package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	middlewares "github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/payments"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

func CreateSessionController(ctx *gin.Context) {
	// Fetching the request body
	var requestBody struct {
		OrderID         *string
		Amount          int
		CustomerDetails types.JuspayCustomer
	}
	ctx.Bind(&requestBody)

	// Creating a Session Controller
	resp, err := middlewares.CreateSession(*requestBody.OrderID, requestBody.Amount, requestBody.CustomerDetails)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.String(http.StatusOK, resp)
}

func CheckOrderStatusController(ctx *gin.Context) {
	// Fetching the request body
	var requestBody struct {
		OrderID *string
	}

	ctx.Bind(&requestBody)
	resp, err := middlewares.CheckOrderStatus(*requestBody.OrderID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.String(http.StatusOK, resp)
}
