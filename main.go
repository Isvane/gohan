package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.PathValue("arg")))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /echo/{arg}", echoHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
