package domain

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Author string `json:"author"`
}

type PostRepository interface {
	Save(post Post) error
	FindOne(id string) (*Post, error)
	FindAll() ([]*Post, error)
}
