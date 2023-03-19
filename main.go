package main

import (
	"olusamimaths/kurunmi/config"
	"olusamimaths/kurunmi/delivery/http/router"
	"olusamimaths/kurunmi/delivery/http/routes"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	run()
}

func run() {
	c := config.NewConfig()
	routes.RegisterRoutes(httpRouter, c)
	httpRouter.SERVE(":8080")
}