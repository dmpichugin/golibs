package valid

import (
	"gopkg.in/go-playground/validator.v9"
)

var (
	v = validator.New()
)

func Struct(s interface{}) error {

	return v.Struct(s)
}
