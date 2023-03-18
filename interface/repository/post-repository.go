package repository

import (
	"log"
	"olusamimaths/kurunmi/domain"
)

type PostRepo struct {
	handler DBHandler
}

func NewPostRepo(handler DBHandler) PostRepo {
	return PostRepo{handler}
}

func (repo PostRepo) Save(post domain.Post) error {
	err := repo.handler.SavePost(&post)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (repo PostRepo) FindAll() ([]*domain.Post, error) {
	posts, err := repo.handler.FindAllPosts()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return posts, nil
}

func (repo PostRepo) FindOne(id string) (*domain.Post, error) {
	result, err := repo.handler.FindPost(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}