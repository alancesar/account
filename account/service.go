package account

type Service interface {
	Create(documentNumber string) (Account, error)
	Get(id uint) (Account, bool, error)
}

func NewAccountService(r Repository) Service {
	return &service{
		r: r,
	}
}

type service struct {
	r Repository
}

func (s *service) Create(documentNumber string) (Account, error) {
	_, exists, err := s.r.GetByDocumentNumber(documentNumber)
	if err != nil {
		return Account{}, err
	} else if exists {
		return Account{}, AlreadyExistsErr
	}

	a := &Account{
		DocumentNumber: documentNumber,
	}

	if err := s.r.Insert(a); err != nil {
		return Account{}, err
	}

	return *a, nil
}

func (s *service) Get(id uint) (Account, bool, error) {
	return s.r.GetByID(id)
}
