package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v", err)
		}
		w, e := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if e != nil {
			fmt.Fprintf(os.Stderr, "fetch:read:%v", e)
		}
		fmt.Printf("%v", w)
	}
}
