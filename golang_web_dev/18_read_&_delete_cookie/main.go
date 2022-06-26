package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/read", read)
	http.HandleFunc("/", dc)

	http.HandleFunc("/delete", delete)
	http.HandleFunc("/create", create)
	http.ListenAndServe(":8082", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("cookie")
	if err != nil {
		cookie = &http.Cookie{
			Name:     "cookie",
			Value:    "This is my first cookie",
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Fprintf(w, "Cookie Created")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("cookie")
	if err != nil {
		http.Redirect(w, r, "/create", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, "%v", c)
}

func delete(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("cookie")
	if err != nil {
		http.Redirect(w, r, "/create", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1 //delete cookie
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func dc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleted")
}
