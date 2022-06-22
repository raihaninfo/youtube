package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/create", create)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func create(w http.ResponseWriter, r *http.Request) {

	cookie, err:= r.Cookie("first-cookie")
	if err!=nil{
		cookie = &http.Cookie{
			Name: "cookie",
			Value: "This is my first cookie",
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}

	fmt.Fprintf(w, "Cookie Created")
}
