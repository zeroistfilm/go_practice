package main

import (
	"fmt"
	"net/http"
)

var c = make(chan *http.Response)

func head(url string) {
	resp, _ := http.Head(url)
	c <- resp
}
func http_url() {

	go head("https://namu.wiki")
	go head("https://www.google.com")
	go head("https://golang.org")

	for i := 0; i < 3; i++ {
		resp := <-c
		if resp != nil {
			resp.Body.Close()
			fmt.Printf("%v: %s\n", resp.Request.URL, resp.Status)
			continue
		}
		fmt.Printf("%v: 연결 실패\n", resp.Request.URL)
	}

}
