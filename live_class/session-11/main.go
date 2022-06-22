package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/asset/", http.StripPrefix("/asset", http.FileServer(http.Dir("./asset"))))
	http.HandleFunc("/", home)
	// http.HandleFunc("/about", about)
	http.HandleFunc("/form", form)
	http.ListenAndServe(":8082", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("home.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

func form(w http.ResponseWriter, r *http.Request) {
	formValue := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(formValue, password)

	// http.Redirect(w, r, "/", http.StatusOK)
}
