package main

import (
	"fmt"
	"log"
	"net/http"
)

func formFunc(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParserForm() err: %v", err)
		return
	}
	fmt.Fprintf(writer, "POST request wass successful\n")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(writer, "Name : %s\n", name)
	fmt.Fprintf(writer, "Address : %s\n", address)
}

func helloFunc(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 Not Found", http.StatusNotFound)
		return
	} else if request.Method != "GET" {
		http.Error(writer, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(writer, "Hello")
}

func main() {
	server := http.FileServer(http.Dir("./static"))
	http.Handle("/", server)
	http.HandleFunc("/hello", helloFunc)
	http.HandleFunc("/form", formFunc)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
