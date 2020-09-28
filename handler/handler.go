package handler

import (
	"errors"
	"github.com/alancesar/account/account"
	"github.com/alancesar/account/transaction"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Message: err.Error(),
	}
}

var InternalServerErr = errors.New("there was an error")

func StartApi(group *gin.RouterGroup, as account.Service, ts transaction.Service) {
	group.POST("accounts", CreateAccount(as))
	group.GET("accounts/:account_id", GetAccount(as))
	group.POST("transactions", CreateTransaction(ts, as))
}

func StartMetrics(group *gin.RouterGroup) {
	group.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"health": true,
		})
	})
}
