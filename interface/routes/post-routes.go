package routes

import (
	"olusamimaths/kurunmi/infrastructure/db"
	"olusamimaths/kurunmi/infrastructure/router"
	"olusamimaths/kurunmi/interface/controllers"
	"olusamimaths/kurunmi/interface/repository"
	"olusamimaths/kurunmi/interface/validators"
	"olusamimaths/kurunmi/usecases"

	"github.com/go-playground/validator/v10"
)

func registerPostRoutes(httpRouter router.Router, postController controllers.PostController) {
	httpRouter.POST("/v1/post", postController.Add)
	httpRouter.GET("/v1/post", postController.FindAll)
	httpRouter.GET("/v1/post/:id", postController.FindOne)
}

func getPostController(dbHandler db.DBHandler) controllers.PostController {
	postRepository := repository.NewPostRepo(dbHandler)
	postInteractor := usecases.NewPostInteractor(postRepository)
	postValidator := validators.NewPostValidator(validator.New())
	return *controllers.NewPostController(postInteractor, postValidator)
}
