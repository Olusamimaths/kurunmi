package main

import (
	"olusamimaths/kurunmi/config"
	"olusamimaths/kurunmi/delivery/http/router"
	"olusamimaths/kurunmi/delivery/http/routes"
	"olusamimaths/kurunmi/infrastructure/db"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	run()
}

func run() {
	c := config.NewConfig()
	dbHandler := db.NewDatabaseHandler(c)

	routes.SetUpDefaultRoutes(httpRouter)

	authorController := routes.GetAuthorController(dbHandler)
	routes.RegisterAuthorRoutes(httpRouter, authorController)

	postController := routes.GetPostController(dbHandler)
	routes.RegisterPostRoutes(httpRouter, postController)

	httpRouter.SERVE(":8080")
}