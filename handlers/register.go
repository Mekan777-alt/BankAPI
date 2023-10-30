package handlers

import (
	"bank/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (h Handler) Registration(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var account models.Account
	err = json.Unmarshal(body, &account)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := "INSERT INTO accounts (first_name, second_name, email, password) VALUES ($1, $2, $3, $4)"

	_, err = h.DB.Exec(query, account.FirstName, account.SecondName, account.Email, account.Password)
	if err != nil {
		log.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}
