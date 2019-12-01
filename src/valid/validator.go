package valid

import (
	"github.com/go-playground/validator"
)

var (
	v = validator.New()
)

func Struct(s interface{}) error {

	return v.Struct(s)
}
