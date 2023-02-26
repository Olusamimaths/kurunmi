package main

import (
	"olusamimaths/kurunmi/src/infrastructure/router"
	"olusamimaths/kurunmi/src/interface/routes"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	routes.RegisterRoutes(httpRouter)
	httpRouter.SERVE(":8080")
}