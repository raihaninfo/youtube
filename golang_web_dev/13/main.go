package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/abc", abc)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		http.NotFoundHandler()
	}
	fmt.Fprintf(w, "Home Page")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func abc(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/about", http.StatusSeeOther)
}
