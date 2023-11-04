package postgres

import "database/sql"

type BillRepository struct {
	DB *sql.DB
}

func NewBillRepository(db *sql.DB) *BillRepository {
	return &BillRepository{db}
}
