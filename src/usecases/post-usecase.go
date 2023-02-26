package usecases

import (
	"log"
	"olusamimaths/kurunmi/src/domain"
)

type PostInteractor struct {
	PostRepository domain.PostRepository
}

func NewPostInteractor(repository domain.PostRepository) PostInteractor {
	return PostInteractor{repository}
}

func (interactor *PostInteractor) CreatePost(post domain.Post) error {
	err := interactor.PostRepository.Save(post)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (interactor *PostInteractor) FindAll() ([]*domain.Post, error) {
	posts, err := interactor.PostRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return posts, nil
}

func (interactor *PostInteractor) FindOne(id string) (*domain.Post, error) {
	result, err := interactor.PostRepository.FindOne(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}