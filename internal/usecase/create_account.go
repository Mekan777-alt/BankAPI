package usecase

import (
	"bank/internal/domain"
	"bank/internal/repository/postgres"
)

type CreateAccountUseCase struct {
	accountRepository postgres.AccountRepository
}

func NewCreateAccountUseCase(accountRepository postgres.AccountRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		accountRepository: accountRepository,
	}
}

type CreateCommand struct {
	FirstName  string
	SecondName string
	Email      string
	Password   string
}

func (useCase *CreateAccountUseCase) Handle(command CreateCommand) (*domain.Account, error) {
	account := domain.NewAccount(command.FirstName, command.SecondName, command.Email, command.Password)

	err := useCase.accountRepository.Save(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
