package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", routeHandler)

	log.Println("shortener started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	route := r.URL.Path
	log.Println("request accepted: ", route)

	var directoryRepo DirectoryRepo
	urls := readDirectoryFromJSON()
	directoryRepo.URLs = urls

	redirectURL, err := directoryRepo.findRedirectURL(route)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "URL is not found!")
	} else {
		log.Println("redirect to: ", redirectURL)
		http.Redirect(w, r, redirectURL, 301)
	}
}
