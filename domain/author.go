package domain

type Author struct {
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
}

type AuthorRepository interface {
	Save(author Author) error
	FindOne(id string) (*Author, error)
	FindAll() ([]*Author, error)
}