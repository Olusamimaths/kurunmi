package controllers

import (
	"encoding/json"
	"net/http"
	"olusamimaths/kurunmi/src/domain"
	"olusamimaths/kurunmi/src/interface/validators"
	"olusamimaths/kurunmi/src/usecases"

	"github.com/go-playground/validator/v10"
)

type AuthorController struct {
	authInteractor usecases.AuthorInteractor
}

func NewAuthorController(interactor usecases.AuthorInteractor) *AuthorController {
	return &AuthorController{interactor}
}

func(controller *AuthorController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var author domain.Author
	err := json.NewDecoder(req.Body).Decode(&author)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid request payload"})
		return
	}

	authorValidator := validators.NewAuthorValidator(validator.New())
	isValid := authorValidator.Validate(author)
	if !isValid {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid author data"})
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