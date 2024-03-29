package db

import (
	"context"
	"log"
	"olusamimaths/kurunmi/domain"

	"go.mongodb.org/mongo-driver/bson"
)

func (dbHandler mongoDBHandler) FindAllPosts() ([]*domain.Post, error) {
	var results []*domain.Post
	collection := dbHandler.database.Collection("posts")

	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem domain.Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		results = append(results, &elem)
	}
	return results, nil
}

func (dbHandler mongoDBHandler) FindPost(id string) (*domain.Post, error) {
	var result *domain.Post
	collection := dbHandler.database.Collection("posts")

	objectId, err := ConvertStringToObjectId(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}

func (dbHandler mongoDBHandler) SavePost(post *domain.Post) error {
	collection := dbHandler.database.Collection("posts")

	_, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
