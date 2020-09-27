package transaction

import (
	"github.com/alancesar/account/account"
	"gorm.io/gorm"
)

type transaction struct {
	gorm.Model
	AccountID     uint
	OperationType string
	Amount        float64
}

func (m *transaction) ToModel(t *Transaction) {
	if t == nil {
		return
	}

	m.AccountID = t.AccountID
	m.OperationType = string(t.OperationType)
	m.Amount = t.Amount
}

func (m *transaction) FromModel() Transaction {
	return Transaction{
		Account: account.Account{
			AccountID: m.AccountID,
		},
		OperationType: OperationType(m.OperationType),
		Amount:        m.Amount,
	}
}
