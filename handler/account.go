package handler

import (
	"fmt"
	"github.com/alancesar/account/account"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createAccountRequest struct {
	DocumentNumber string `json:"document_number" binding:"required"`
}

type getAccountRequest struct {
	AccountID uint `uri:"account_id" binding:"required"`
}

func CreateAccount(s account.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req createAccountRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, NewErrorResponse(err))
			return
		}

		a, err := s.Create(req.DocumentNumber)
		if err != nil {
			if err == account.AlreadyExistsErr {
				msg := fmt.Errorf("%s already has an account", req.DocumentNumber)
				context.JSON(http.StatusBadRequest, NewErrorResponse(msg))
				return
			}

			context.JSON(http.StatusInternalServerError, NewErrorResponse(InternalServerErr))
			return
		}

		context.JSON(http.StatusCreated, gin.H{
			"account_id":      a.AccountID,
			"document_number": a.DocumentNumber,
		})
	}
}

func GetAccount(s account.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req getAccountRequest
		if err := context.ShouldBindUri(&req); err != nil {
			context.JSON(http.StatusBadRequest, NewErrorResponse(err))
			return
		}

		a, exists, err := s.Get(req.AccountID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, NewErrorResponse(InternalServerErr))
			return
		} else if !exists {
			msg := fmt.Errorf("account %d does not exists", req.AccountID)
			context.JSON(http.StatusNotFound, NewErrorResponse(msg))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"account_id":      a.AccountID,
			"document_number": a.DocumentNumber,
		})
	}
}
