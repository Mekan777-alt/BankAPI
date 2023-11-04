package handlers

import (
	"bank/internal/domain"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var JsonResponse domain.Account
	err = json.Unmarshal(body, &JsonResponse)
	if err != nil {
		log.Fatal(err)
	}
	var acount domain.Account
	query := "SELECT email, password FROM accounts WHERE email = $1"
	row := h.DB.QueryRow(query, JsonResponse.Email)
	err = row.Scan(&acount.Email, &acount.Password)
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
	if acount.Password == JsonResponse.Password {
		w.Header().Add("Content-Type", "applicatin/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Entry completed")

	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Not fount Login or password")
	}
}
