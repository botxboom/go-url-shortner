package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/shorten", shortenURL)
	// http.HandleFunc("/:shortURL", redirectURL)

	http.ListenAndServe(":8080", nil)

}
