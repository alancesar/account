package transaction

import "gorm.io/gorm"

type Repository interface {
	Insert(transaction *Transaction) error
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&transaction{})
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (s *repository) Insert(t *Transaction) error {
	var m transaction
	m.ToModel(t)

	if query := s.db.Create(&m); query.Error != nil {
		return query.Error
	}

	t.TransactionID = m.ID
	return nil
}

