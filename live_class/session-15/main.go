package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/login", login)
	r.HandleFunc("/loginauth", loginauth)
	r.HandleFunc("/profile", profile)
	r.HandleFunc("/logout", logout)
	http.ListenAndServe(":8082", r)
}

var store = sessions.NewCookieStore([]byte("secret-key"))

func home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Println(err)
	}
	temp.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "login-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, ok := session.Values["username"]

	if !ok {
		temp, err := template.ParseFiles("templates/login.html")
		if err != nil {
			fmt.Println(err)
		}
		temp.Execute(w, nil)

	}
	if ok {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}

}

func loginauth(w http.ResponseWriter, r *http.Request) {

	dummyUser := "admin"
	dummyPassword := "12345"

	userName := r.FormValue("username")
	password := r.FormValue("password")

	if userName == dummyUser && password == dummyPassword {
		// Login code
		session, err := store.Get(r, "login-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 1,
			HttpOnly: true,
		}
		session.Values["username"] = "username"
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	} else {
		// Invalid
		temp, err := template.ParseFiles("templates/login.html")
		if err != nil {
			fmt.Println(err)
		}
		temp.Execute(w, "Please Check your username and password")
	}

}

func profile(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "login-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, ok := session.Values["username"]

	if ok {
		temp, err := template.ParseFiles("templates/profile.html")
		if err != nil {
			fmt.Println(err)
		}
		temp.Execute(w, nil)
	}
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	delete(session.Values, "username")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
