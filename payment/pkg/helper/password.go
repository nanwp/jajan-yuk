package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func GeneratePasswordString(bahan string) (string, error) {
	hasher := md5.New()
	hasher.Write([]byte(bahan))
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
