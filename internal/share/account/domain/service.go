package domain

import "go.mongodb.org/mongo-driver/mongo"

type CheckingAccountValidityService struct {
	accountRepository Repository
}

func NewCheckingAccountValidityService(accountRepository Repository) *CheckingAccountValidityService {
	return &CheckingAccountValidityService{
		accountRepository: accountRepository,
	}
}

func (ds *CheckingAccountValidityService) IsValidated(account *Account) bool {
	// ID唯一性校验
	_, err := ds.accountRepository.GetByID(account.ID)
	return err == mongo.ErrNoDocuments
}
