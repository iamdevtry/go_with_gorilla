package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	router := mux.NewRouter()
	port := ":8080"
	router.HandleFunc("/hello", HealthHandler).Methods("GET")
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Println("Server is running on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
