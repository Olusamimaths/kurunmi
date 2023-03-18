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

func registerAuthorRoutes(httpRouter router.Router, authorController controllers.AuthorController) {
	httpRouter.POST("/v1/author", authorController.Add)
	httpRouter.GET("/v1/author", authorController.FindAll)
	httpRouter.GET("/v1/author/:id", authorController.FindOne)
}

func getAuthorController(dbHandler db.DBHandler) controllers.AuthorController {
	authorRepository := repository.NewAuthorRepo(dbHandler)
	authorInteractor := usecases.NewAuthorInteractor(authorRepository)
	authorValidator := validators.NewAuthorValidator(validator.New())
	return *controllers.NewAuthorController(authorInteractor, authorValidator)
}