package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}
	h3 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #3!\n")
	}
	h4 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a awesome HandleFunc #4!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/h2", h2)
	http.HandleFunc("/h3", h3)
	http.HandleFunc("/h4", h4)

	fmt.Println("Server starting")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
