package main

import (
	"olusamimaths/kurunmi/src/config"
	"olusamimaths/kurunmi/src/infrastructure/router"
	"olusamimaths/kurunmi/src/interface/routes"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	c := config.NewConfig()
	routes.RegisterRoutes(httpRouter, c)
	httpRouter.SERVE(":8080")
}