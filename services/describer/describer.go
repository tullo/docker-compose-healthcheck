package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var descriptions = []string{
	"awesome 🤩",
	"goofy 🙃",
	"creative 🤔",
	"an awesome/goofy/creative developer 🐳",
}

func main() {
	// Simulate a non-trivial service startup time.
	time.Sleep(10 * time.Second)

	fmt.Println("Describer is good to go!")

	http.ListenAndServe(":http", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(descriptions[rand.Intn(len(descriptions))]))
	}))
}
