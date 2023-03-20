package controllers

import (
	"encoding/json"
	"net/http"
	"olusamimaths/kurunmi/domain"
	"olusamimaths/kurunmi/interface/messages"
	"olusamimaths/kurunmi/interface/validators"
	"olusamimaths/kurunmi/usecases"

	"github.com/gorilla/mux"
)

type PostController struct {
	postInteractor usecases.PostInteractor
	validator validators.PostValidator
}

func NewPostController(interactor usecases.PostInteractor, validator validators.PostValidator) *PostController {
	return &PostController{interactor, validator }
}

func(controller *PostController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var post domain.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: messages.InvalidRequestPayload})
		return
	}

	isValid := controller.validator.Validate(post)
	if !isValid {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: messages.InvalidPostData})
		return
	}

	err = controller.postInteractor.CreatePost(post)
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

	vars := mux.Vars(req)
	id := vars["id"]
	if id == "" {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(ErrorResponse{Message: messages.InvalidId})
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