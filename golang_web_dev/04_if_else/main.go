package main

import (
	"net/http"
	"text/template"
)

type Member struct {
	Name     string
	Language string
	Member   bool
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8082", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	member:= Member{
		Name: "Raihan",
		Language: "???",
		Member: false,
	}

	tpl.Execute(w, member)
}
