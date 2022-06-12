package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/abc", abc)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page"))
}
func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}
func abc(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/about", http.StatusSeeOther)
}
