package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	url := "http://placekitten.com/g/640/340"

	res, err := http.Get(url)
	defer func() {
		err = res.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	file, err := os.Create("sample.jpg")
	if err != nil {
		panic(err)
	}
	io.Copy(file, res.Body)
}
