package main

import (
	"log"
	"net/http"

	"github.com/devnandito/imageapi/api"
	"github.com/gorilla/mux"
)

func main() {
	addr := ":8080"
	router := mux.NewRouter()
	router.HandleFunc("/api/image", api.HandleApiShowImage).Methods("GET")
	router.HandleFunc("/api/image/{id}", api.HandleApiGetOneImage).Methods("GET")
	router.HandleFunc("/api/image", api.HandleApiCreateImage).Methods("POST")
	log.Printf("Server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}