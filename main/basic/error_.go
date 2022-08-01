package basic

import (
	"log"
	"os"
)

func Error_() {
	f, err := os.Open("./go.mod")
	if err != nil {
		log.Fatal(err)
	}
	println(f.Name())
}
