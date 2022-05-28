package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir("./style")))
	http.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/home", home)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("home.html", "header.html", "footer.html")
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
