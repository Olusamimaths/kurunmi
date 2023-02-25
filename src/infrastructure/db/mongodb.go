package db

import (
	"context"
	"log"
	"olusamimaths/kurunmi/src/domain"
	"olusamimaths/kurunmi/src/interface/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoDBHandler interface {
	repository.DBHandler
}
type mongoDBHandler struct {
	MongoClient mongo.Client
	database *mongo.Database

}

func NewMongoDBHandler(connectString string, dbname string) (IMongoDBHandler, error) {
	dbHandler := mongoDBHandler{}
	clientOptions := options.Client().ApplyURI(connectString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}

	err =  client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}

	dbHandler.MongoClient = *client
	dbHandler.database = client.Database(dbname)
	return dbHandler, nil
}

func (dbHandler mongoDBHandler) FindAllPosts() ([]*domain.Post, error) {
	var results []*domain.Post
	collection := dbHandler.database.Collection("posts")

	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem domain.Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		results = append(results, &elem)
	}
	return results, nil
}

func (dbHandler mongoDBHandler) FindPost(id string) (*domain.Post, error) {
	var result *domain.Post
	collection := dbHandler.database.Collection("posts")

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func (dbHandler mongoDBHandler) SavePost(post *domain.Post) error {
	collection := dbHandler.database.Collection("posts")

	_, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (dbHandler mongoDBHandler) FindAllAuthors() ([]*domain.Author, error) {
	var results []*domain.Author
	collection := dbHandler.database.Collection("authors")

	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem domain.Author
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		results = append(results, &elem)
	}
	return results, nil
}


func (dbHandler mongoDBHandler) FindAuthor(id string) (*domain.Author, error) {
	var result *domain.Author
	collection := dbHandler.database.Collection("authors")

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func (dbHandler mongoDBHandler) SaveAuthor(author *domain.Author) error {
	collection := dbHandler.database.Collection("authors")

	_, err := collection.InsertOne(context.TODO(), author)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}


