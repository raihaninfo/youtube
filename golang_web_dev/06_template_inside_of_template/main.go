package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("index.html", "header.html", "footer.html")
	if err != nil {
		log.Println(err)
	}
	tem.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("about.html", "header.html", "footer.html")
	if err != nil {
		log.Println(err)
	}
	tem.Execute(w, nil)
}
