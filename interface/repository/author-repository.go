package repository

import (
	"log"
	"olusamimaths/kurunmi/domain"
)

type AuthorRepo struct {
	handler DBHandler
}

func NewAuthorRepo(handler DBHandler) AuthorRepo {
	return AuthorRepo{handler}
}

func (repo AuthorRepo) Save(author domain.Author) error {
	err := repo.handler.SaveAuthor(&author)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (repo AuthorRepo) FindAll() ([]*domain.Author, error) {
	authors, err := repo.handler.FindAllAuthors()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return authors, nil
}

func (repo AuthorRepo) FindOne(id string) (*domain.Author, error) {
	result, err := repo.handler.FindAuthor(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}