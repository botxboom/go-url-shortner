package handler

import (
	"net/http"
)

func shortenURL(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	r.ParseForm()

	// Get the original URL
	original := r.Form.Get("url")

	// Validate the URL
	if original == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// // Create a new URL
	// url := models.URL{
	// 	Original: original,
	// }

	// Save the URL
	// err := url.Save()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Return the shortened URL
	w.Write([]byte(original))
}
