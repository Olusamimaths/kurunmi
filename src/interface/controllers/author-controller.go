package controllers

import (
	"encoding/json"
	"net/http"
	"olusamimaths/kurunmi/src/domain"
	"olusamimaths/kurunmi/src/interface/messages"
	"olusamimaths/kurunmi/src/interface/validators"
	"olusamimaths/kurunmi/src/usecases"
)

type AuthorController struct {
	authInteractor usecases.AuthorInteractor
	validator validators.AuthorValidator
}

func NewAuthorController(interactor usecases.AuthorInteractor, validator validators.AuthorValidator) *AuthorController {
	return &AuthorController{interactor, validator}
}

func(controller *AuthorController) Add(res http.ResponseWriter, req *http.Request) {
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

func(controller *AuthorController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	authors, err := controller.authInteractor.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(res).Encode(authors)
}