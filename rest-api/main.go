package main

import (
	"log"
	"net/http"

	"github.com/Sreejith1962/rest-api/db"
	"github.com/Sreejith1962/rest-api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	collection := db.ConnectDb()

	handler := &handlers.Handler{Collection: collection}

	router.HandleFunc("/books", handler.GetBooks).Methods("GET")
	router.HandleFunc("/books", handler.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handler.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", handler.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handler.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
