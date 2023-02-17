package main

import (
	"fmt"

	"github.com/swapnilwalke/Cryptography/otp"
)

func main() {
	message := otp.Otp("Hello")
	fmt.Println(message)
}
