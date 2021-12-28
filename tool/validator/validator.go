package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phone", phoneValidator)
	}
}

func phoneValidator(f validator.FieldLevel) bool {
	data := f.Field().String()
	return len(data) == 11
}
