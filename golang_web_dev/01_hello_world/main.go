package main

import (
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("Welcome to Learn With Raihan")
	http.HandleFunc("/", home)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello World!")
}