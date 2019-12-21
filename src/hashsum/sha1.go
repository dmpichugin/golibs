package hashsum

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func Sha1(data string) string {

	hash := sha1.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func Sha1Reader(reader io.Reader) (string, error) {

	hasher := sha1.New()

	_, err := io.Copy(hasher, reader)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
