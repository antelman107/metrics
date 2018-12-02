package nosql

import (
	"crypto/rand"
	"encoding/base64"
)

func getRandString() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	val := "'" + base64.StdEncoding.EncodeToString(b) + "'"
	return val, nil
}
