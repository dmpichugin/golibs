package uniq

import (
	uuid "github.com/satori/go.uuid"
)

func UUID4() string {

	return uuid.NewV4().String()
}

func UUID5(nameSpace string) string {

	return uuid.NewV5(uuid.NewV4(), nameSpace).String()
}
