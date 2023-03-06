package db

import (
	"context"
	"log"
	"olusamimaths/kurunmi/src/domain"

	"go.mongodb.org/mongo-driver/bson"
)

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