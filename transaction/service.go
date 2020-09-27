package transaction

import "github.com/alancesar/account/account"

type Service interface {
	Create(account account.Account, amount float64, ot OperationType) (Transaction, error)
}

func NewTransactionService(r Repository) Service {
	return &service{
		r: r,
	}
}

type service struct {
	r Repository
}

func (s *service) Create(account account.Account, amount float64, ot OperationType) (Transaction, error) {
	a, err := ParseAmount(amount, ot)
	if err != nil {
		return Transaction{}, err
	}

	t := &Transaction{
		Account:       account,
		OperationType: ot,
		Amount:        a,
	}

	if err := s.r.Insert(t); err != nil {
		return Transaction{}, err
	}

	return *t, nil
}
