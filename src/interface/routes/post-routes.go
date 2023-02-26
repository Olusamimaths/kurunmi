package routes

import (
	"olusamimaths/kurunmi/src/infrastructure/router"
	"olusamimaths/kurunmi/src/interface/controllers"
	"olusamimaths/kurunmi/src/interface/repository"
	"olusamimaths/kurunmi/src/interface/validators"
	"olusamimaths/kurunmi/src/usecases"

	"github.com/go-playground/validator/v10"
)

func registerPostRoutes(httpRouter router.Router, postController controllers.PostController) {
	httpRouter.POST("/post/add", postController.Add)
	httpRouter.GET("/post", postController.FindAll)
	httpRouter.GET("/post/:id", postController.FindOne)
}

func getPostController() controllers.PostController {
	postRepository := repository.NewPostRepo(dbHandler)
	postInteractor := usecases.NewPostInteractor(postRepository)
	postValidator := validators.NewPostValidator(validator.New())
	return *controllers.NewPostController(postInteractor, postValidator)
}
