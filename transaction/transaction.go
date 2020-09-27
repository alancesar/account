package transaction

import "github.com/alancesar/account/account"

type OperationType string

const (
	ForCash     OperationType = "FOR_CACHE"
	Installment OperationType = "INSTALLMENT"
	Withdrawn   OperationType = "WITHDRAWN"
	Payment     OperationType = "PAYMENT"
)

type Transaction struct {
	TransactionID uint
	account.Account
	OperationType OperationType
	Amount        float64
}

var (
	OperationTypes = map[string]OperationType{
		string(ForCash):     ForCash,
		string(Installment): Installment,
		string(Withdrawn):   Withdrawn,
		string(Payment):     Payment,
	}
)
