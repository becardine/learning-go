package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// File handling in Go

	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	// writeString adicionar texto ao arquivo, usado quando sabe que o texto que será adicionado é string, exemplo:
	// size, err := f.WriteString("Hello, World!")

	// write adiciona bytes ao arquivo, usado quando não sabe o tipo de dado que será adicionado, exemplo:
	size, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote %d bytes\n", size)

	// fecha o arquivo após a escrita para liberar recursos
	f.Close()

	// leitura de arquivo
	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	// converte o arquivo em bytes para string
	fmt.Println(string(file))

	// imagina a situação onde o arquivo é muito grande e não é possível ler tudo de uma vez
	// para isso, é possível ler o arquivo por partes
	// primeiro faz a leitura aos poucos, de acordo com o recurso disponível e depois abre o arquivo
	file2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	// cria um buffer para armazenar os bytes lidos
	reader := bufio.NewReader(file2)

	// lê 5 bytes por vez
	buffer := make([]byte, 5)

	// enquanto houver bytes para ler, continua lendo
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		// converte os bytes lidos em string
		fmt.Println(string(buffer[:n]))
	}
	file2.Close()

	// deleta o arquivo
	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}

}
