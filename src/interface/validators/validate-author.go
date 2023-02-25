package validators

import (
	"fmt"
	"olusamimaths/kurunmi/src/domain"

	"github.com/go-playground/validator/v10"
)

type AuthorValidator interface {
	Validate(author domain.Author) bool

	ValidateName(title string) bool
	ValidateUsername(body string) bool
	ValidateEmail(author string) bool
}

type authorValidator struct {
	validator *validator.Validate
}

func (aV *authorValidator) ValidateEmail(email string) bool {
	errs := aV.validator.Var(email, "required, email")
	if errs != nil {
		fmt.Println(errs)
		return false
	}
	return true
}

func (aV *authorValidator) ValidateName(name string) bool {
	errs := aV.validator.Var(name, "required, min=3, max=20")
	if errs != nil {
		fmt.Println(errs)
		return false
	}
	return true
}

func (aV *authorValidator) ValidateUsername(username string) bool {
	errs := aV.validator.Var(username, "required, min=3, max=20")
	if errs != nil {
		fmt.Println(errs)
		return false
	}
	return true
}

func NewAuthorValidator(val *validator.Validate) AuthorValidator {
	return &authorValidator{validator: val}
}

func (pV *authorValidator) Validate(author domain.Author) bool {
	return pV.ValidateName(author.Name) &&
		pV.ValidateUsername(author.Username) &&
		pV.ValidateEmail(author.Email)
}
