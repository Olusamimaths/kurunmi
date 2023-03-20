package routes

import (
	"olusamimaths/kurunmi/delivery/http/router"
	"olusamimaths/kurunmi/interface/controllers"
)

func RegisterAuthorRoutes(httpRouter router.Router, authorController controllers.AuthorController) {
	httpRouter.POST("/v1/author", authorController.Add)
	httpRouter.GET("/v1/author", authorController.FindAll)
	httpRouter.GET("/v1/author/{id}", authorController.FindOne)
}

