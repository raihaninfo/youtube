package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	// var s string
	if r.Method == http.MethodPost {
		f, _, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
		}

		// 234234.png
		// file234532.png
		tempFile, err := ioutil.TempFile("upload", "image*.png")
		if err != nil {
			panic(err)
		}

		defer tempFile.Close()
		tempFile.Write(bs)

		// s = string(bs)
		// fmt.Println(s)

	}
	temp, err := template.ParseFiles("form.html")
	if err != nil {
		log.Println(err)
	}
	temp.Execute(w, nil)
}
