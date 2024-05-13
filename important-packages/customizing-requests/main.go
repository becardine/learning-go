package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://httpbin.org/delay/2", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
