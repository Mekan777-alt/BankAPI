package models

import "time"

type Card struct {
	ID             uint      `json:"ID" gorm:"primaryKey"`
	Number         string    `json:"number"`
	Cvv            string    `json:"cvv"`
	ExpirationDate time.Time `json:"expirationDate"`
	Balance        float64   `json:"balance"`
	History        []History `json:"history" gorm:"foreignKey:ID"`
	IsCardActive   bool      `json:"isCardActive"`
}
