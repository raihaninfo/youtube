package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8082", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	// var s string
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
		}
		contentType := h.Header["Content-Type"][0]
		fmt.Println(contentType)

		var tempFile *os.File
		if contentType == "image/png" {
			tempFile, err = ioutil.TempFile("upload/png", "image*.png")
			if err != nil {
				panic(err)
			}
		}else if contentType == "image/jpeg"{
			tempFile, err = ioutil.TempFile("upload/jpg", "image*.jpg")
			if err != nil {
				panic(err)
			}
		}

		// 234234.png
		// file234532.png

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
