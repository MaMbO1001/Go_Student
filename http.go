package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Задание 1.

type Hello struct {
	Message string `json:"message"`
}

func jopa(w http.ResponseWriter, r *http.Request) {
	hello := Hello{
		Message: "Hello world",
	}
	jsonData, err := json.Marshal(hello)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// Задание 2.

type GGWP struct {
	Messag string `json:"messag"`
}

func jopa1(w http.ResponseWriter, r *http.Request) {
	ggwp := GGWP{
		Messag: "Что-то написано",
	}
	jsonData, err := json.Marshal(ggwp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

type GG struct {
	Messag string `json:"messag"`
}

func jopa2(w http.ResponseWriter, r *http.Request) {
	gg := GG{
		Messag: "Добро пожаловать домой",
	}
	jsonData, err := json.Marshal(gg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// Задание 3.

func jopa3(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

// Задание 4.


func main() {

	// Задание 1.
	http.Handle("/", http.HandlerFunc(jopa))

	// Задание 2.

	http.Handle("/about", http.HandlerFunc(jopa1))
	http.Handle("/home", http.HandlerFunc(jopa2))

	// Задание 3.

	http.Handle("/greet", http.HandlerFunc(jopa3))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
