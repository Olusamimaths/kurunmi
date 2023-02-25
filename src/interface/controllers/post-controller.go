package controllers

import (
	"encoding/json"
	"net/http"
	"olusamimaths/kurunmi/src/domain"
	"olusamimaths/kurunmi/src/interface/validators"
	"olusamimaths/kurunmi/src/usecases"

	"github.com/go-playground/validator/v10"
)

type PostController struct {
	postInteractor usecases.PostInteractor
	validator validators.PostValidator
}

func NewBookController(interactor usecases.PostInteractor, validator validators.PostValidator) *PostController {
	return &PostController{interactor, validator }
}

func(controller *PostController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var post domain.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid request payload"})
		return
	}

	postValidator := validators.NewPostValidator(validator.New())
	isValid := postValidator.Validate(post)
	if !isValid {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid post data"})
		return
	}

	err = controller.postInteractor.CreateBook(post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusCreated)
}

func(controller *PostController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	posts, err := controller.postInteractor.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func(controller *PostController) FindOne(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	id := req.URL.Query().Get("id")
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid post id"})
		return
	}

	post, err := controller.postInteractor.FindOne(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(post)
}