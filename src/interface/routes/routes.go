package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"olusamimaths/kurunmi/src/infrastructure/db"
	"olusamimaths/kurunmi/src/infrastructure/router"
	"olusamimaths/kurunmi/src/interface/repository"
)

var (
	dbHandler repository.DBHandler
)

func setUpDefaultRoutes(httpRouter router.Router) {
	httpRouter.GET("/api", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{"message": "Ok", "status": "success"})
	})
}

func RegisterRoutes(httpRouter router.Router) {
	dbConnected := connectToDatabase()
	if !dbConnected {
		return
	}
	setUpDefaultRoutes(httpRouter)

	authorController := getAuthorController()
	registerAuthorRoutes(httpRouter, authorController)

	postController := getPostController()
	registerPostRoutes(httpRouter, postController)
}

func connectToDatabase() bool {
	var err error
	dbHandler, err = db.NewDBHandler("mongodb://localhost:27017", "kurunmi")
	if err != nil {
		log.Println("Unable to connect to database")
		log.Fatal(err.Error())
		return false
	}
	return true
}
