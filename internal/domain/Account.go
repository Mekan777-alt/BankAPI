package domain

import "github.com/gofrs/uuid"

type Account struct {
	id         uuid.UUID `json:"ID"`
	firstName  string    `json:"first_name"`
	secondName string    `json:"second_name"`
	email      string    `json:"email"`
	password   string    `json:"password"`
	bill       []Bill    `json:"bills"`
	status     bool      `json:"status"`
}

func NewAccount(firstname, secondname, email, password string) *Account {
	return &Account{
		id:         uuid.Must(uuid.NewV7()),
		firstName:  firstname,
		secondName: secondname,
		email:      email,
		password:   password,
		status:     true,
	}
}

func (a *Account) ID() uuid.UUID      { return a.id }
func (a *Account) FirstName() string  { return a.firstName }
func (a *Account) SecondName() string { return a.secondName }
func (a *Account) Email() string      { return a.email }
func (a *Account) Password() string   { return a.password }
func (a *Account) Status() bool       { return a.status }

type AccountRepository interface {
}
