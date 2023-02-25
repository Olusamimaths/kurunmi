package usecases

import (
	"log"
	"olusamimaths/kurunmi/src/domain"
)

type AuthorInteractor struct {
	AuthorRepository domain.AuthorRepository
}

func NewAuthorInteractor(repository domain.AuthorRepository) AuthorInteractor {
	return AuthorInteractor{repository}
}

func (interactor *AuthorInteractor) CreateAuthor(Author domain.Author) error {
	err := interactor.AuthorRepository.Save(Author)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (interactor *AuthorInteractor) FindAll() ([]*domain.Author, error) {
	Authors, err := interactor.AuthorRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return Authors, nil
}

func (interactor *AuthorInteractor) FindOne(id string) (*domain.Author, error) {
	result, err := interactor.AuthorRepository.FindOne(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}