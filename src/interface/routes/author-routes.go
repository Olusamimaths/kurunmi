package routes

import (
	"olusamimaths/kurunmi/src/infrastructure/router"
	"olusamimaths/kurunmi/src/interface/controllers"
	"olusamimaths/kurunmi/src/interface/repository"
	"olusamimaths/kurunmi/src/interface/validators"
	"olusamimaths/kurunmi/src/usecases"

	"github.com/go-playground/validator/v10"
)

func registerAuthorRoutes(httpRouter router.Router, authorController controllers.AuthorController) {
	httpRouter.POST("/author/add", authorController.Add)
	httpRouter.GET("/author", authorController.FindAll)
	httpRouter.GET("/author/:id", authorController.FindOne)
}

func getAuthorController() controllers.AuthorController {
	authorRepository := repository.NewAuthorRepo(dbHandler)
	authorInteractor := usecases.NewAuthorInteractor(authorRepository)
	authorValidator := validators.NewAuthorValidator(validator.New())
	return *controllers.NewAuthorController(authorInteractor, authorValidator)
}