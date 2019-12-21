package hashsum

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func MD5(data string) string {

	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func MD5Reader(reader io.Reader) (string, error) {

	hasher := md5.New()

	_, err := io.Copy(hasher, reader)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
