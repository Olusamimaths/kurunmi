package repository

import "olusamimaths/kurunmi/src/domain"

type DBHandler interface {
	FindAllPosts() ([]*domain.Post, error)
	FindPost(id string) (*domain.Post, error)
	SavePost(post *domain.Post) error

	FindAllAuthors() ([]*domain.Author, error)
	FindAuthor(id string) (*domain.Author, error)
	SaveAuthor(author *domain.Author) error
}