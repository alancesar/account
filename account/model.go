package account

import "gorm.io/gorm"

type account struct {
	gorm.Model
	DocumentNumber string
}

func (m *account) ToModel(a *Account) {
	if a == nil {
		return
	}

	m.DocumentNumber = a.DocumentNumber
}

func (m *account) FromModel() Account {
	return Account{
		AccountID:      m.ID,
		DocumentNumber: m.DocumentNumber,
	}
}
