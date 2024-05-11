package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	request, err := http.Get("https://golang.org")
	if err != nil {
		panic(err)
	}

	// defer é usado para adiar a execução de uma função até o final do bloco atual
	// no caso, a função request.Body.Close() será executada após a execução do bloco main
	defer request.Body.Close()

	res, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))

	// exemplo de uso do defer
	fmt.Println("first")        // será executado primeiro
	defer fmt.Println("second") // será executado após o bloco main
	fmt.Println("third")        // será executado após o primeiro print
}
