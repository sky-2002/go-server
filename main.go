package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful !")
	name := r.Form.Get("name")
	address := r.Form.Get("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func helloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello from Golang server 😇!!")
}

func main() {
	// we plan to build three routes
	// which are '/' , '/hello', and '/form'
	// and corressponding html files and functions like
	// index.html, hello function, and form function with form.html

	fileServer := http.FileServer(http.Dir("./static")) // this will check for index.html file in static directory

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandlerFunc)

	http.HandleFunc("/hello", helloHandlerFunc)

	fmt.Println("Starting golang server on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

// func helloHandlerFunc() string {
// 	return "Hello from golang !!"
// }
