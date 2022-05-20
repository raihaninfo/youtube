package main

import (
	
	"net/http"
	"text/template"
)

var temp *template.Template
var err error

func main() {
	temp, err = template.ParseGlob("template/*.html")
	if err!=nil{
		panic(err)
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	// temp, err := template.ParseFiles("template/index.html")
	// if err != nil {
	// 	log.Println(err)
	// }
	// temp.Execute(w, nil)
	temp.ExecuteTemplate(w, "index.html", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	// temp, err := template.ParseFiles("template/index.html")
	// if err != nil {
	// 	log.Println(err)
	// }
	// temp.Execute(w, nil)
	temp.ExecuteTemplate(w, "about.html", nil)
}
