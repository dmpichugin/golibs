package uniq

import (
	"github.com/gofrs/uuid"
	uuid2 "github.com/google/uuid"
)

func UUIDv4() string {

	u, err := uuid.NewV4()
	if err != nil {
		return uuid2.New().String()
	}

	return u.String()
}
