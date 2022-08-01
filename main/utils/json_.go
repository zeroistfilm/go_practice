package utils

import (
	"encoding/json"
	"fmt"
)

func Json_() {

	jsonEncoding()
}

type Member struct {
	Name   string
	Age    int
	Active bool
}

func jsonEncoding() {
	mem := Member{"Alex", 10, true}
	jsonBytes, err := json.Marshal(mem)
	fmt.Println(jsonBytes)
	if err != nil {
		panic(err)
	}
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
}
