package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		value := r.Form.Get("cycle")
		log.Print(value)
		cycle, err := strconv.Atoi(value)
		if err != nil {
			log.Print(err)
		}
		lissajous(w, float64(cycle))
	}
	http.HandleFunc("/lissajous", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
