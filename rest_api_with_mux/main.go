package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func main() {
	router := mux.NewRouter()
	users = append(users, User{ID: "1", Name: "Raihan", Email: "raiha@gmail.com", Password: "password"})

	router.HandleFunc("/api/user", getUser).Methods("GET")
	router.HandleFunc("/api/user", createUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", getUserById).Methods("GET")
	router.HandleFunc("/api/user/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", deleteUser).Methods("DELETE")

	log.Println("Listening port on localhost:8080")
	http.ListenAndServe(":8080", router)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)
	for _, value := range users {
		if value.ID == id["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)
	for index, item := range users {
		if item.ID == id["id"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			json.NewDecoder(r.Body).Decode(&user)
			users = append(users, user)
			json.NewEncoder(w).Encode(&user)
			return
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)
	for index, item := range users {
		if item.ID == id["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
}
