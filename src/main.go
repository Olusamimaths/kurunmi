package main

import (
	"fmt"
	"log"
	"net/http"
	"olusamimaths/kurunmi/src/infrastructure/db"
	"olusamimaths/kurunmi/src/infrastructure/router"
	"olusamimaths/kurunmi/src/interface/controllers"
	"olusamimaths/kurunmi/src/interface/repository"
	"olusamimaths/kurunmi/src/interface/validators"
	"olusamimaths/kurunmi/src/usecases"

	"github.com/go-playground/validator/v10"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	dbHandler db.IMongoDBHandler
)

func getPostController() controllers.PostController {
	postRepository := repository.NewPostRepo(dbHandler)
	postInteractor := usecases.NewPostInteractor(postRepository)
	postValidator := validators.NewPostValidator(validator.New())
	return *controllers.NewPostController(postInteractor, postValidator)
}

func getAuthorController() controllers.AuthorController {
	authorRepository := repository.NewAuthorRepo(dbHandler)
	authorInteractor := usecases.NewAuthorInteractor(authorRepository)
	authorValidator := validators.NewAuthorValidator(validator.New())
	return *controllers.NewAuthorController(authorInteractor, authorValidator)
}

func main() {
		httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "App is up and running...")
	})

	var err error
	dbHandler, err = db.NewMongoDBHandler("mongodb://localhost:27017", "kurunmi")
	if err != nil {
		log.Println("Unable to connect to database")
		log.Fatal(err.Error())
		return
	}

	postController := getPostController()
	authorController := getAuthorController()

	httpRouter.POST("/post/add", postController.Add)
	httpRouter.GET("/post", postController.FindAll)

	httpRouter.POST("/author/add", authorController.Add)

	httpRouter.SERVE(":8080")
}