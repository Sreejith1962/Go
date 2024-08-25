package routes

import (
	"github.com/Sreejith1962/go-bookstore/pkg/controller"
	"github.com/gorilla/mux"
)

// var RegisterBookStoreRoutes = func(router *mux.Router)
func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookid}", controller.GetBooksById).Methods("GET")
	router.HandleFunc("/book/{bookid}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controller.DeleteBook).Methods("DELETE")

}
