package postgres

import (
	"bank/internal/domain"
	"database/sql"
)

type AccountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (a *AccountRepository) Save(account *domain.Account) error {
	query := "INSERT INTO accounts (first_name, second_name, email, password) VALUES ($1, $2, $3, $4)"

	_, err := a.DB.Exec(query, account.FirstName(), account.SecondName(), account.Email(), account.Password())
	if err != nil {
		return err
	}
	return nil
}
