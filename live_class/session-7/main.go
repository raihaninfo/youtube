package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var temp *template.Template
var err error

type UserInfo struct {
	Name   string
	Age    int
	Member bool
	// MemberInfo Member
}

// type Member struct {
// 	TotalMember  int
// 	ActiveMember bool
// }

func main() {
	fmt.Println("Welcome to Learn With Raihan")
	temp, err = template.ParseGlob("templates/*.html")
	HandleErr(err)
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "home.html", nil)
	HandleErr(err)
}

func about(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("templates/about.html")
	HandleErr(err)
	// age := 21

	name := "Rezaul"
	age := 21
	member := true

	user := UserInfo{
		Name:   name,
		Age:    age,
		Member: member,
	}

	tem.Execute(w, user)
}

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}
