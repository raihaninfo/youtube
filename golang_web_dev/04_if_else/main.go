package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8082", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	tpl.Execute(w, nil)
}
