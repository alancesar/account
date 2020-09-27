package account

import "gorm.io/gorm"

type Repository interface {
	Insert(account *Account) error
	GetByID(id uint) (Account, bool, error)
	GetByDocumentNumber(documentNumber string) (Account, bool, error)
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&account{})
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Insert(a *Account) error {
	var m account
	m.ToModel(a)

	if query := r.db.Create(&m); query.Error != nil {
		return query.Error
	}

	a.AccountID = m.ID
	return nil
}

func (r *repository) GetByID(id uint) (Account, bool, error) {
	var m account

	if query := r.db.First(&m, id); query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return Account{}, false, nil
		}

		return Account{}, false, query.Error
	}

	return m.FromModel(), true, nil
}

func (r *repository) GetByDocumentNumber(documentNumber string) (Account, bool, error) {
	var m account

	if query := r.db.Where("document_number = ?", documentNumber).First(&m); query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return Account{}, false, nil
		}

		return Account{}, false, query.Error
	}

	return m.FromModel(), true, nil
}
