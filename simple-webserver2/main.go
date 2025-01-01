package main

import (
	"log"
	"net/http"
)

func main() {
	// FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root.
	http.Handle("/", http.FileServer(http.Dir("static")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start:", err)
	}
}
