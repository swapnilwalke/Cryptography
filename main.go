package main

import (
	"log"
	"os"
)

// ! https://learn.gopherguides.com/courses/preparing-your-environment-for-go-development/modules/setting-up-windows/#slide-3
func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("Your GOPATH has not been set !!")
	}
	log.Println(gopath)
	log.Println("Success!")
}
