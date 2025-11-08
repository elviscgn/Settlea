package uid

import (
	"crypto/rand"
	"errors"
	"log"
)

var defaultAlphabet = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const defaultIDSize = 12

func MustGenerateID(size int) (string, error) {
	if size <= 0 {
		return "", errors.New("size must be a positive integer")
	}

	id := make([]rune, size)
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i := 0; i < size; i++ {
		id[i] = defaultAlphabet[bytes[i]%byte(len(defaultAlphabet))]
	}
	return string(id), nil
}

func GenerateUniqueID(size int) string {
	id, err := MustGenerateID(size)
	if err != nil {
		log.Fatal(err)
	}

	return id

}
