package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"olusamimaths/kurunmi/domain"
	"olusamimaths/kurunmi/interface/messages"
	"olusamimaths/kurunmi/interface/validators"
	"olusamimaths/kurunmi/usecases"

	"github.com/gorilla/mux"
)

type AuthorController struct {
	authInteractor usecases.AuthorInteractor
	validator      validators.AuthorValidator
}

func NewAuthorController(interactor usecases.AuthorInteractor, validator validators.AuthorValidator) *AuthorController {
	return &AuthorController{interactor, validator}
}

func (controller *AuthorController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var author domain.Author
	err := json.NewDecoder(req.Body).Decode(&author)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: messages.InvalidRequestPayload})
		return
	}

	isValid := controller.validator.Validate(author)
	if !isValid {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: messages.InvalidAuthorData})
		return
	}

	err = controller.authInteractor.CreateAuthor(author)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusCreated)
}

func (controller *AuthorController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	authors, err := controller.authInteractor.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(res).Encode(authors)
}

func (controller *AuthorController) FindOne(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	fmt.Println("FindOne called")

	vars := mux.Vars(req)
	id := vars["id"]
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: messages.InvalidId})
		return
	}

	post, err := controller.authInteractor.FindOne(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(post)
}
