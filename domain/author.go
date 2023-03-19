package domain

type Author struct {
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type AuthorRepository interface {
	Save(author Author) error
	FindOne(id string) (*Author, error)
	FindAll() ([]*Author, error)
}