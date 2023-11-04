package handlers

import (
	"bank/internal/domain"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func randomNumber(numb int) string {
	var number string
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < numb; i++ {
		number += strconv.Itoa(r.Intn(10))
	}
	return number
}

func (h Handler) CreateCard(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var JsonResponse domain.Bill
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&JsonResponse)
	if err != nil {
		log.Println(err)
	}
	queryBillId := "SELECT id FROM bills WHERE number = $1"
	bill_id := h.DB.QueryRow(queryBillId, JsonResponse.Number)

	card := domain.Card{
		Number:         randomNumber(16),
		Cvv:            randomNumber(3),
		ExpirationDate: time.Now().AddDate(5, 0, 0),
		Balance:        0,
		History:        nil,
		IsCardActive:   true,
	}
	query := "INSERT INTO cards (bill_id, number, cvv, expiration_date, iscardactive, balance) VALUES ($1, $2, $3, $4, $5)"
	_, err = h.DB.Exec(query, *bill_id, card.Number, card.Cvv, card.ExpirationDate, card.IsCardActive, card.Balance)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)

}
