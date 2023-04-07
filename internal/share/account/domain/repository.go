package domain

type Repository interface {
	GetAll() ([]*Account, error)
	GetByID(accountID string) (*Account, error)
	Save(account *Account) error
	Replace(account *Account) error
	Delete(accountID string) error
}
