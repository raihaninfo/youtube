package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/form", form)
	http.HandleFunc("/thanks", thanks)
	http.ListenAndServe(":8082", nil)
}

func form(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("form.html")
	if err != nil {
		log.Println(err)
	}
	temp.Execute(w, nil)
}

func thanks(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("thanks.html")
	if err != nil {
		log.Println(err)
	}
	temp.Execute(w, nil)
}
