package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get_() {
	customGet()
}

func normalGet() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("%s\n", string(data))
}

func customGet() {
	req, err := http.NewRequest("GET", "http://csharp.tips/feed/rss", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Crawler")

	client := http.Client{} //...?????
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	fmt.Println(str)

}
