package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	if r.Method == http.MethodPost {
		file, h, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		bs, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		contentType := h.Header["Content-Type"][0]

		fmt.Println(contentType)

		var tempFile *os.File

		if contentType == "image/png" {
			tempFile, err = ioutil.TempFile("upload/png/", "image*.png")
			if err != nil {
				panic(err)
			}
		} else if contentType == "image/jpeg" {
			tempFile, err = ioutil.TempFile("upload/jpg/", "image*.jpg")
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println("File Extension not allowed")
		}

		tempFile.Write(bs)

	}

	temp, err := template.ParseFiles("filevalue.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)

}
