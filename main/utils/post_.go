package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Post_() {
	//simplePost()
	//postForm()
	jsonPost()
}

func simplePost() {
	reqBody := bytes.NewBufferString("Post plain text")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", reqBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	str := string(respBody)
	println(str)
}

func postForm() {
	reqBody := url.Values{"Name": {"Lee"}, "Age": {"10"}}
	resp, err := http.PostForm("http://httpbin.org/post", reqBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	str := string(respBody)
	println(str)
}

type Person struct {
	Name string
	Age  int
}

func jsonPost() {
	person := Person{"Alex", 10}
	fmt.Println(person)
	pbytes, err := json.Marshal(person)
	fmt.Println(pbytes)
	buff := bytes.NewBuffer(pbytes)
	fmt.Println(buff)
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	str := string(respBody)
	println(str)
}
