package domain

type CheckingAccountValidityService struct {
	accountRepository Repository
}

func NewCheckingAccountValidityService(accountRepository Repository) *CheckingAccountValidityService {
	return &CheckingAccountValidityService{
		accountRepository: accountRepository,
	}
}

func (ds *CheckingAccountValidityService) IsValidated(account Account) bool {
	// ID唯一性校验
	result, _ := ds.accountRepository.Get(account.ID)
	return result == nil
}
