package main

import (
	"fmt"
	"log"
	"net/http"
)

func amazHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Amazon Customer !!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/amaz", amazHandler)

	fmt.Println("Server started at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Hello!")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, "parseform error", err)
		return
	}
	fmt.Fprintln(w, "Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintln(w, "Name: ", name)
	fmt.Fprintln(w, "Address: ", address)

}
