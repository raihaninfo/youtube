package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/asset/", http.StripPrefix("/asset", http.FileServer(http.Dir("./asset"))))
	http.HandleFunc("/", home)
	// http.HandleFunc("/about", about)
	http.HandleFunc("/form", form)
	http.ListenAndServe(":8082", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("home.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

func form(w http.ResponseWriter, r *http.Request) {

	var bsString string
	if r.Method == http.MethodPost {
		file, _, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		bs, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		bsString = string(bs)
	}

	temp, err := template.ParseFiles("filevalue.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, bsString)

}
