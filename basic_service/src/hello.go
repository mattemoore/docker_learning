package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", os.Getenv("HOSTNAME"))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready.")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ready", readyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
