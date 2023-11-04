package domain

import (
	"github.com/gofrs/uuid"
	"math/rand"
	"strconv"
	"time"
)

type Bill struct {
	id           uuid.UUID   `json:"ID"`
	number       string      `json:"number"`
	limit        int         `json:"limit"`
	cards        []uuid.UUID `json:"cards"`
	isBillActive bool        `json:"isBillActive"`
}

func randomNumberBill() string {
	var number string
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < 20; i++ {
		number += strconv.Itoa(r.Intn(10))
	}
	return number
}

func NewBill() *Bill {
	return &Bill{
		id:           uuid.Must(uuid.NewV7()),
		number:       randomNumberBill(),
		limit:        0,
		cards:        nil,
		isBillActive: true,
	}
}

func (b *Bill) Number() string     { return b.number }
func (b *Bill) Limit() int         { return b.limit }
func (b *Bill) Cards() []uuid.UUID { return b.cards }
func (b *Bill) IsBillActive() bool { return b.isBillActive }

func (b *Bill) Close() {
	b.isBillActive = false
}
