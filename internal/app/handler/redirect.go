package handler

import (
	"net/http"
)

func redirectURL(w http.ResponseWriter, r *http.Request) {

	// shortURL := r.URL.Path[1:]

	// url, err := models.GetURL(shortURL)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	// http.Redirect(w, r, url.Original, http.StatusMovedPermanently)
}
