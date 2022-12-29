package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func generatePassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	password := generateRandomString(10)
	fmt.Println(password)
	pass := ""
	hashedPassword := generatePassword(pass)
	fmt.Println(hashedPassword)
}
