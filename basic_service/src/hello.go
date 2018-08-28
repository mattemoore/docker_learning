package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World.")
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready.")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ready", readyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
