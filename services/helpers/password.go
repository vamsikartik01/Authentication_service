package helpers

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"log"
)

func GenerateSalt(password string) (string, error) {
	saltSize := 10
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt[:])
	if err != nil {
		log.Println("Error Generating Salt")
		return string(salt), err
	}
	return base64.URLEncoding.EncodeToString(salt), nil
}

func GeneratePasswordHash(password string, salt string) (string, error) {
	passwordSalt := password + salt
	passwordBytes := []byte(passwordSalt)
	hasher := sha512.New()
	hasher.Write(passwordBytes)
	passwordHashedBytes := hasher.Sum(nil)
	return base64.URLEncoding.EncodeToString(passwordHashedBytes), nil
}

func GetSaltAndHash(password string) (string, string, error) {
	salt, err := GenerateSalt(password)
	if err != nil {
		return "", "", err
	}
	passwordHashed, err := GeneratePasswordHash(password, salt)
	if err != nil {
		return "", "", err
	}
	return passwordHashed, salt, nil
}

func VerifySaltAndPassword(password string, salt string, hash string) (bool, error) {
	passwordHashed, err := GeneratePasswordHash(password, salt)
	if err != nil {
		return false, err
	}
	if passwordHashed == hash {
		return true, nil
	}
	return false, nil
}
