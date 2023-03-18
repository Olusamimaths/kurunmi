package validators

import (
	"fmt"
	"olusamimaths/kurunmi/domain"

	"github.com/go-playground/validator/v10"
)

type PostValidator interface {
	Validate(post domain.Post) bool
	
	ValidateTitle(title string) bool
	ValidateBody(body string) bool
	ValidateAuthor(author string) bool
}

type postValidator struct {
	validator *validator.Validate
}

func NewPostValidator(val *validator.Validate) PostValidator {
	return &postValidator{validator: val}
}

func (pV *postValidator) Validate(post domain.Post) bool {
	return pV.ValidateTitle(post.Title) &&
		pV.ValidateBody(post.Body) &&
		pV.ValidateAuthor(post.Author)
}

func (pV *postValidator) ValidateAuthor(title string) bool {
	errs := pV.validator.Var(title, "required,min=3,max=20")
	if errs != nil {
		fmt.Println(errs)
		return false
	}
	return true
}

func (pV *postValidator) ValidateBody(body string) bool {
	errs := pV.validator.Var(body, "required,min=10")
	if errs != nil {
		fmt.Println(errs)
		return false
	}
	return true
}

func (pV *postValidator) ValidateTitle(author string) bool {
	errs := pV.validator.Var(author, "required,min=3,max=100")
	if errs != nil {
		fmt.Println(errs)
		return false
	}
	return true
}
