package main

import (
	"net/http"
	"strconv"
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
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	if idNumb != 1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Mario not found"))
		return
	}
	w.Write([]byte("Hello, Mario!"))
}
