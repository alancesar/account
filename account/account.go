package account

import "errors"

type Account struct {
	AccountID      uint
	DocumentNumber string
}

var AlreadyExistsErr = errors.New("account already exists")