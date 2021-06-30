package main

import (
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, 5)
	}
	http.HandleFunc("/lissajous", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
