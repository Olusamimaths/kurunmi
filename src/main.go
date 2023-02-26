package main

import (
	"olusamimaths/kurunmi/src/infrastructure/db"
	"olusamimaths/kurunmi/src/infrastructure/router"
	"olusamimaths/kurunmi/src/interface/controllers"
	"olusamimaths/kurunmi/src/interface/repository"
	"olusamimaths/kurunmi/src/interface/validators"
	"olusamimaths/kurunmi/src/usecases"

	"github.com/go-playground/validator/v10"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	dbHandler db.IMongoDBHandler
)

func getPostController() controllers.PostController {
	postRepository := repository.NewPostRepo(dbHandler)
	postInteractor := usecases.NewPostInteractor(postRepository)
	postValidator := validators.NewPostValidator(validator.New())
	return *controllers.NewPostController(postInteractor, postValidator)
}

func getAuthorController() controllers.AuthorController {
	authorRepository := repository.NewAuthorRepo(dbHandler)
	authorInteractor := usecases.NewAuthorInteractor(authorRepository)
	authorValidator := validators.NewAuthorValidator(validator.New())
	return *controllers.NewAuthorController(authorInteractor, authorValidator)
}