package main

import (
	"mvc/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.Home)
	http.ListenAndServe(":8080", nil)
}
