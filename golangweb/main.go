package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/hello", http.HandlerFunc(helloHandlet))

	http.ListenAndServe(":8080", mux)
}

func helloHandlet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
