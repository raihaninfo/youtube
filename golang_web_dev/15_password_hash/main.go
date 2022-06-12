package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password123"
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Main Password", password)
	fmt.Println("Hash Password", string(pass))

	password1 := "password"
	compare := bcrypt.CompareHashAndPassword([]byte(pass), []byte(password1))
	if compare == nil {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}
