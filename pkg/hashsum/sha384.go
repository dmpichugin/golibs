package hashsum

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
)

func Sha384(data string) string {

	hash := sha512.Sum384([]byte(data))
	return hex.EncodeToString(hash[:])
}

func Sha384Reader(reader io.Reader) (string, error) {

	hasher := sha512.New384()

	_, err := io.Copy(hasher, reader)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
