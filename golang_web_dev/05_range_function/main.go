package main

import (
	"net/http"
	"text/template"
)

// type GroceryList []string
type Todo struct {
	Name string
	Done bool
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

	// grocerylist := GroceryList{"Milk", "Eggs", "Sugar", "Tea"}
	todos := []Todo{{"Picup Groceries", true}, {"Pain Kitchen", false}, {"Make Video", true}, {"Upload Video", false}}

	tpl.Execute(w, todos)
}
