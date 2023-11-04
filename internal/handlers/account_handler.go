package handlers

import (
	"bank/internal/usecase"
	"encoding/json"
	"github.com/gofrs/uuid"
	"net/http"
)

type POSTAccountsHandler struct {
	useCase *usecase.CreateAccountUseCase
}

func NewPOSTAccountsHandler(useCase *usecase.CreateAccountUseCase) *POSTAccountsHandler {
	return &POSTAccountsHandler{useCase: useCase}
}

type POSTAccountRequest struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type POSTAccountResponse struct {
	id     uuid.UUID
	status bool
}

func (responce *POSTAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id     uuid.UUID `json:"id"`
		Status bool      `json:"isOpened"`
	}{
		Id:     responce.id,
		Status: responce.status,
	})
}

func (handler *POSTAccountsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body POSTAccountRequest
	err := json.NewDecoder(request.Body).Decode(&body)

	command := usecase.CreateCommand{
		FirstName:  body.FirstName,
		SecondName: body.SecondName,
		Email:      body.Email,
		Password:   body.Password,
	}

	account, err := handler.useCase.Handle(command)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	response := &POSTAccountResponse{
		id:     account.ID(),
		status: !account.Status(),
	}

	writer.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}
}
