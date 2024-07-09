package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func main() {
	key, err := generateRandomString(32)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}
	fmt.Println("Generated JWT Secret Key:", key)
}
