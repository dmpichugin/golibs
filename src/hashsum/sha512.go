package hashsum

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
)

func Sha512(data string) string {

	hash := sha512.Sum512([]byte(data))
	return hex.EncodeToString(hash[:])
}

func Sha512Reader(reader io.Reader) (string, error) {

	hasher := sha512.New()

	_, err := io.Copy(hasher, reader)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
