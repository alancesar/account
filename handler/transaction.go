package handler

import (
	"errors"
	"fmt"
	"github.com/alancesar/account/account"
	"github.com/alancesar/account/transaction"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createTransactionRequest struct {
	AccountID     uint    `json:"account_id" binding:"required"`
	OperationType string  `json:"operation_type" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
}

func CreateTransaction(ts transaction.Service, as account.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req createTransactionRequest
		if err := context.BindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, NewErrorResponse(err))
			return
		}

		a, exists, err := as.Get(req.AccountID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, NewErrorResponse(InternalServerErr))
			return
		} else if !exists {
			msg := fmt.Errorf("account %d doesn't exist", req.AccountID)
			context.JSON(http.StatusBadRequest, NewErrorResponse(msg))
			return
		}

		ot, exists := transaction.OperationTypes[req.OperationType]
		if !exists {
			msg := fmt.Errorf("%s is an invalid opetation type", req.OperationType)
			context.JSON(http.StatusBadRequest, NewErrorResponse(msg))
			return
		}

		t, err := ts.Create(a, req.Amount, ot)
		if err != nil {
			context.JSON(http.StatusInternalServerError, NewErrorResponse(errors.New("there was an error")))
			return
		}

		context.JSON(http.StatusCreated, gin.H{
			"transaction_id": t.TransactionID,
			"account_id":     t.AccountID,
			"operation_type": t.OperationType,
			"amount":         t.Amount,
		})
	}
}
