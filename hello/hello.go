package main

import (
	"fmt"

	"github.com/swapnilwalke/Cryptography/otp"
)

func main() {
	//message := otp.Otp("Hello")
	//fmt.Println(message)
	err := otp.Otp("swapnil")
	if err != "" {
		fmt.Println(err)
	}

	terr := otp.GenerateKey(9)
	fmt.Print(terr)

}
