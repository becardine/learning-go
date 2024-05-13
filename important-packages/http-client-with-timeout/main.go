package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	// Timeout é um recurso importante para evitar que a aplicação fique esperando por muito tempo por uma resposta de uma requisição
	// O timeout é definido em segundos
	// primeiro é criado um client com um timeout de 1 segundo
	// client := http.Client{Timeout: time.Second}
	// depois é feita uma requisição para o endereço https://httpbin.org/delay/2
	// resp, err := client.Get("https://httpbin.org/delay/2")
	// o tempo de resposta é de 2 segundos, então a requisição vai falhar
	// if err != nil {
	// o erro é impresso na tela, que é um timeout
	// Get "https://httpbin.org/delay/2": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
	// panic(err)
	// }
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(body))

	// criar um post com timeout
	client := http.Client{}
	json := bytes.NewBuffer([]byte(`{"name": "golang"}`))
	resp, err := client.Post("https://httpbin.org/delay/2", "application/json", json)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
