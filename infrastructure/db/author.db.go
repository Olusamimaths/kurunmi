package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"olusamimaths/kurunmi/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (dbHandler mongoDBHandler) FindAllAuthors() ([]*domain.Author, error) {
	var results []*domain.Author
	collection := dbHandler.database.Collection("authors")

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cur, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem domain.Author
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		results = append(results, &elem)
	}
	return results, nil
}

func (dbHandler mongoDBHandler) FindAuthor(id string) (*domain.Author, error) {
	var result *domain.Author
	collection := dbHandler.database.Collection("authors")

	objectId, err := ConvertStringToObjectId(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: objectId}}

	opts := options.FindOne().SetProjection(bson.D{{Key: "password", Value: 0}})
	err = collection.FindOne(context.TODO(), filter, opts).Decode(&result)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}

func (dbHandler mongoDBHandler) SaveAuthor(author *domain.Author) error {
	collection := dbHandler.database.Collection("authors")

	filter := bson.D{{Key: "email", Value: author.Email}}
	err := checkForExistingEntity(collection, filter)
	fmt.Println("error: ", err)
	if err != nil {
		return errors.New("email already exists")
	}

	filter = bson.D{{Key: "username", Value: author.Username}}
	err = checkForExistingEntity(collection, filter)
	if err != nil {
		return errors.New("username already exists")
	}

	_, err = collection.InsertOne(context.TODO(), author)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func checkForExistingEntity(collection *mongo.Collection, filter primitive.D) (error) {
	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == nil {
		if result != nil {
			return errors.New("entity already exists")
		}
	}
	return nil
}
