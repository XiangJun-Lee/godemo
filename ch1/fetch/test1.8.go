package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		fmt.Println(url)
		fmt.Println(strings.HasPrefix(url, prefix))
		if !strings.HasPrefix(url, "http://") {
			fmt.Println(url + " not have http head!")
			// 字符串拼接
			url = prefix + url
		}
		fmt.Println("new url=" + url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v", err)
		}
		b, e := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if e != nil {
			fmt.Fprintf(os.Stderr, "fetch:read:%v", e)
		}
		fmt.Printf("%s", b)
	}
}
