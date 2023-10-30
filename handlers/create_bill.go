package handlers

import (
	"bank/models"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func randomNumberBill() string {
	var number string
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < 20; i++ {
		number += strconv.Itoa(r.Intn(10))
	}
	return number
}

func (h Handler) CreateBill(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var users models.Account
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&users)
	query := "SELECT * FROM accounts WHERE email = $1"
	row := h.DB.QueryRow(query, users.Email)

	var dbUser models.Account
	err = row.Scan(&dbUser.ID, &dbUser.FirstName, &dbUser.SecondName, &dbUser.Email, &dbUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if users.Password != dbUser.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		var bill = models.Bill{
			ID:           dbUser.ID,
			Number:       randomNumberBill(),
			Limit:        0,
			Cards:        nil,
			IsBillActive: true,
		}

		query = "INSERT INTO bills (account_id, number, card, sum_limit) VALUES ($1, $2, $3, $4)"
		_, err = h.DB.Exec(query, bill.ID, bill.Number, nil, bill.Limit)
		if err != nil {
			log.Println(err)
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bill)
	}
}
