package routes

import (
	"olusamimaths/kurunmi/delivery/http/router"
	"olusamimaths/kurunmi/interface/controllers"
)

func RegisterPostRoutes(httpRouter router.Router, postController controllers.PostController) {
	httpRouter.POST("/v1/post", postController.Add)
	httpRouter.GET("/v1/post", postController.FindAll)
	httpRouter.GET("/v1/post/:id", postController.FindOne)
}


