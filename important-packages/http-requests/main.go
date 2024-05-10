package main

import (
	"io"
	"net/http"
)

func main() {

	// requeste para o site do golang
	request, err := http.Get("https://golang.org")
	if err != nil {
		panic(err)
	}

	// lê o corpo da resposta
	res, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	// converte os bytes lidos em string
	println(string(res))

	// fecha o corpo da resposta após a leitura
	// é importante fechar o corpo da resposta para liberar recursos
	request.Body.Close()
}
