package hashsum

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func Sha256(data string) string {

	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func Sha256Reader(reader io.Reader) (string, error) {

	hasher := sha256.New()

	_, err := io.Copy(hasher, reader)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
