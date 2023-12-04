package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Задание 4.

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	Idd := mux.Vars(r)
	id, _ := strconv.Atoi(Idd["id"])
	for _, user := range Users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = len(Users) + 1
	Users = append(Users, user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	Idd := mux.Vars(r)
	id, _ := strconv.Atoi(Idd["id"])
	for index, user := range Users {
		if user.ID == id {
			Users[index].Name = Idd["name"]
			json.NewEncoder(w).Encode(Users[index])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	Idd := mux.Vars(r)
	id, _ := strconv.Atoi(Idd["id"])
	for index, user := range Users {
		if user.ID == id {
			Users = append(Users[:index], Users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Users)
}

func main() {

	fmt.Println("Hello")
	// Задание 4.
	router := mux.NewRouter()

	Users = append(Users, User{ID: 412, Name: "Viktor"})
	Users = append(Users, User{ID: 121, Name: "Edgar"})

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}/{name}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
    
	log.Fatal(http.ListenAndServe(":8080", router))
}

