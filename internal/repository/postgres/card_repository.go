package postgres

import "database/sql"

type CardRepository struct {
	DB *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{db}
}
