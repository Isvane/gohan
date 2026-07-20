package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"sync"
)

type Database struct {
	mu       sync.RWMutex
	UserInfo map[string]int
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.PathValue("arg")))
}

func main() {
	db := &Database{
		UserInfo: make(map[string]int),
	}
	_ = db

	mux := http.NewServeMux()

	mux.HandleFunc("GET /echo/{arg}", echoHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
