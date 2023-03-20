package db

import (
	"context"
	"errors"
	"log"
	"olusamimaths/kurunmi/config"
	"olusamimaths/kurunmi/interface/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBHandler interface {
	repository.DBHandler
}

type mongoDBHandler struct {
	MongoClient mongo.Client
	database    *mongo.Database
}


func NewDatabaseHandler(c *config.Config) DBHandler {
	dbHandler, err := getDbHandler(c)

	if err != nil {
		log.Fatal(err.Error())
		panic("Unable to connect to database")
	}
	return dbHandler
}

func getDbHandler(c *config.Config) (DBHandler, error) {
	connectString := c.Get().GetString("db.url")
	dbname := c.Get().GetString("db.name")
	dbHandler := mongoDBHandler{}
	clientOptions := options.Client().ApplyURI(connectString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}

	dbHandler.MongoClient = *client
	dbHandler.database = client.Database(dbname)
	return dbHandler, nil
}

func ConvertStringToObjectId(id string) (primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return objectId, errors.New("invalid id")
	}
	return objectId, nil
}