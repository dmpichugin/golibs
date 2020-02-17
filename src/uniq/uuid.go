package uniq

import (
	"github.com/gofrs/uuid"
)

func UUID() string {

	return uuid.Must(uuid.NewV4()).String()
}
