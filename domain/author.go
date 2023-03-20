package domain

type Author struct {
	Id string `json:"id" bson:"_id,omitempty"` 
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password,omitempty"`
}

type AuthorRepository interface {
	Save(author Author) error
	FindOne(id string) (*Author, error)
	FindAll() ([]*Author, error)
}