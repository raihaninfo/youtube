package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name  string
	Price int
	Qt    int
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

	// name:= "Raihan"

	product := Product{
		Name:  "Apple",
		Price: 21,
		Qt:    5,
	}

	tpl.Execute(w, product)
}
