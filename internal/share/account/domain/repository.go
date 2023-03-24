package domain

type Repository interface {
	Get(accountID AccountId) (*Account, error)
	GetAll() ([]Account, error)
	Update(account Account) error
	Save(account Account) error
	Delete(accountID AccountId) error
}
