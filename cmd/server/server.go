package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	fmt.Fprintln(w, "Server Response")
}

func main() {
	log.Println("Starting server")
	http.HandleFunc("/", handler)
	log.Println("Listening on Port 8081")
	http.ListenAndServe(":8081", nil)
}
