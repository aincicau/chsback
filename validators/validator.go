package validators

import (
	"github.com/go-playground/validator"
)

var validate *validator.Validate

func SetValidator() {
	validate = validator.New()
}

func GetValidator() *validator.Validate {
	return validate
}
