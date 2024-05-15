package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// cria um contexto vazio
	ctx := context.Background()

	// cria um contexto com cancelamento automático após 1 segundo
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	// por boas práticas, sempre chame a função cancel ao final do uso do contexto
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
