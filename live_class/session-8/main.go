package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var temp *template.Template
var err error

func main() {
	fmt.Println("Welcome to Learn With Raihan")
	temp, err = template.ParseGlob("templates/*.html")
	HandleErr(err)
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "home.html", nil)
	HandleErr(err)
}

func about(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("templates/about.html")
	HandleErr(err)

	// student:= []string{"Rezaul", "Tanvir", "Sabbir", "Ekram"}
	country := map[string]string{
		"Dhaka":        "Bangladesh",
		"Delhi":        "India",
		"Kuyalalampur": "Malayasia",
	}

	tem.Execute(w, country)
}

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}
