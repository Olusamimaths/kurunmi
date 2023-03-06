package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"olusamimaths/kurunmi/src/config"
	"olusamimaths/kurunmi/src/infrastructure/db"
	"olusamimaths/kurunmi/src/infrastructure/router"
)



func setUpDefaultRoutes(httpRouter router.Router) {
	httpRouter.GET("/api", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{"message": "Ok", "status": "success"})
	})

	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{"message": "App is running...", "status": "success"})
	})
}

func RegisterRoutes(httpRouter router.Router, c *config.Config) {
	dbHandler := db.NewDatabaseHandler(c)

	setUpDefaultRoutes(httpRouter)

	authorController := getAuthorController(dbHandler)
	registerAuthorRoutes(httpRouter, authorController)

	postController := getPostController(dbHandler)
	registerPostRoutes(httpRouter, postController)
	fmt.Println("Routes set up successfully")
}

