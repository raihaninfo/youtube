package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Listening Posrt :8082")
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Method)
	w.Write([]byte("This Is my home page"))
}
