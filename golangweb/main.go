package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/hello", http.HandlerFunc(helloHandler))
	mux.Handle("/mario", http.HandlerFunc(helloMarioHandler))

	http.ListenAndServe(":8080", mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func helloMarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Mario!"))
}
