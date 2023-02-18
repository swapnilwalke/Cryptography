package otp

import (
	"fmt"
	"math/rand"
)

func Otp(text string) string {
	fmt.Println(text)
	return text
}

func GenerateKey(length int) []byte {
	// !! make initiates slice, map or channel of specific length in Go
	key := make([]byte, length)
	rand.Read(key)
	return key
}
