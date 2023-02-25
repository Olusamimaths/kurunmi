package repository

import (
	"log"
	"olusamimaths/kurunmi/src/domain"
)

type AuthorRepo struct {
	handler DBHandler
}

func NewAuthorRepo(handler DBHandler) AuthorRepo {
	return AuthorRepo{handler}
}

func (repo *AuthorRepo) SaveAuthor(author domain.Author) error {
	err := repo.handler.SaveAuthor(&author)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (repo *AuthorRepo) FindAllAuthors() ([]*domain.Author, error) {
	authors, err := repo.handler.FindAllAuthors()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return authors, nil
}

func (repo *AuthorRepo) FindOneAuthor(id string) (*domain.Author, error) {
	result, err := repo.handler.FindAuthor(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}