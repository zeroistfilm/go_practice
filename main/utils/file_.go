package utils

import (
	"io"
	"os"
)

func Files_() {
	fi, err := os.Open("./main/main.go")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fo, err := os.Create("./example.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024)

	for {
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if cnt == 0 {
			break
		}
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}

}
