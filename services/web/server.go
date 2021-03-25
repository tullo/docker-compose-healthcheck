package main

import (
	"fmt"
	"io"
	"net/http"
)

func getDescription() (string, error) {
	response, err := http.Get("http://describer/")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	description, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(description), nil
}

func main() {
	// Grab an initial description just to test that we can access the describer
	// service. If this fails, then exit immediately.
	// This isn't necessary, but it's useful to simulate a hard dependency on
	// another service that needs to be running and healthy before this service
	// can succeed.
	if _, err := getDescription(); err != nil {
		panic("unable to grab initial description")
	}

	fmt.Println("Web is good to go!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		description, err := getDescription()
		if err != nil {
			http.Error(w, "unable to grab description", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf(
			"<!doctype html><html><body><h1>You are %s</h1></body></html>",
			description,
		)))
	})

	http.ListenAndServe(":http", nil)
}
