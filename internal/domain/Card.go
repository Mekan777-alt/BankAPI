package domain

import (
	"github.com/gofrs/uuid"
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	id             uuid.UUID `json:"ID"`
	number         string    `json:"number"`
	cvv            string    `json:"cvv"`
	expirationDate time.Time `json:"expirationDate"`
	balance        float64   `json:"balance"`
	isCardActive   bool      `json:"isCardActive"`
}

func randomNumber(numb int) string {
	var number string
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < numb; i++ {
		number += strconv.Itoa(r.Intn(10))
	}
	return number
}

func NewCard() *Card {
	return &Card{
		id:             uuid.Must(uuid.NewV7()),
		number:         randomNumber(16),
		cvv:            randomNumber(3),
		expirationDate: time.Now().AddDate(5, 0, 0),
		balance:        0,
		isCardActive:   true,
	}
}

func (c *Card) Number() string            { return c.number }
func (c *Card) Cvv() string               { return c.cvv }
func (c *Card) ExpirationDate() time.Time { return c.expirationDate }
func (c *Card) Balance() float64          { return c.balance }
func (c *Card) IsCardActive() bool        { return c.isCardActive }

func (c *Card) Close() {
	c.isCardActive = false
}
