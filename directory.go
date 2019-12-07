package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

// DirectoryRepo :nodoc:
type DirectoryRepo struct {
	URLs URLs
}

// URLs :nodoc:
type URLs struct {
	URLs []URL `json:"urls"`
}

// URL :nodoc:
type URL struct {
	BaseURL     string `json:"base_url"`
	RedirectURL string `json:"redirect_url"`
}

func (r *DirectoryRepo) findRedirectURL(baseURL string) (redirectURL string, err error) {
	for _, url := range r.URLs.URLs {
		if url.BaseURL == baseURL {
			redirectURL = url.RedirectURL
		}
	}

	if redirectURL == "" {
		return redirectURL, errors.New("url not found")
	}

	return
}

func readDirectoryFromJSON() (urls URLs) {
	jsonFile, err := os.Open("urls.json")
	if err != nil {
		log.Println(err)
		return
	}
	defer jsonFile.Close()

	jsonByte, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(jsonByte, &urls)

	return
}
