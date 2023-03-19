package routes

import (
	"encoding/json"
	"net/http"
	"olusamimaths/kurunmi/delivery/http/router"
	"olusamimaths/kurunmi/infrastructure/db"
	"olusamimaths/kurunmi/interface/controllers"
	"olusamimaths/kurunmi/interface/repository"
	"olusamimaths/kurunmi/interface/validators"
	"olusamimaths/kurunmi/usecases"

	"github.com/go-playground/validator/v10"
)



func SetUpDefaultRoutes(httpRouter router.Router) {
	httpRouter.GET("/api", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{"message": "Ok", "status": "success"})
	})

	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{"message": "App is running...", "status": "success"})
	})
}


func GetAuthorController(dbHandler db.DBHandler) controllers.AuthorController {
	authorRepository := repository.NewAuthorRepo(dbHandler)
	authorInteractor := usecases.NewAuthorInteractor(authorRepository)
	authorValidator := validators.NewAuthorValidator(validator.New())
	return *controllers.NewAuthorController(authorInteractor, authorValidator)
}

func GetPostController(dbHandler db.DBHandler) controllers.PostController {
	postRepository := repository.NewPostRepo(dbHandler)
	postInteractor := usecases.NewPostInteractor(postRepository)
	postValidator := validators.NewPostValidator(validator.New())
	return *controllers.NewPostController(postInteractor, postValidator)
}