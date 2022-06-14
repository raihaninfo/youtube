package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Learn With Raihan")

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// w.Write([]byte("<h2>This is home page</h2>"))
	fmt.Fprintf(w, "<h2 style='color:green;'>This is Home page</h2>")
}

func about(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	// w.Write([]byte("<h2>This is home page</h2>"))
	fmt.Fprintf(w, "<h2 style='color:green;'>This is About page</h2>")
	
}