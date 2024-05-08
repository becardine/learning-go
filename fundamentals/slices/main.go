package main

import "fmt"

func main() {
	// Declarando um slice de 5 posições
	// O slice é uma referência a um array. Dividido em 3 partes: ponteiro, tamanho e capacidade
	// O tamanho é para ele saber onde ir, e capacidade é para ele saber o quanto ele consegue receber de informação
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("length: %d cap=%d %v\n", len(s), cap(s), s)

	// : é um operador de slice, que permite pegar um pedaço do slice. é um slice de slice
	// tudo que está antes do 0, vai ser desconsiderado
	// apesar de ficar 0 itens, a capacidade ainda é 5
	// a saída será []
	fmt.Printf("length: %d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// tudo que está antes do 4, vai ser desconsiderado
	// a capacidade ainda é 5
	// a saída será [1 2 3 4]
	fmt.Printf("length: %d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// vai ser considerado tudo que está depois do 2
	// a capacidade será 3
	// a saída será [3 4 5]
	fmt.Printf("length: %d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// adicionando um item ao slice
	// toda vez que o slice é alterado, caso ele não tenha mais capacidade, ele cria um novo array com o dobro da capacidade
	// a saída será [1 2 3 4 5 6]
	// dica: sempre que possível, defina a capacidade do slice o mais próximo do que você acha que vai ser necessário, para evitar que o slice seja recriado várias vezes
	s = append(s, 16)
	fmt.Printf("length: %d cap=%d %v\n", len(s[:1]), cap(s[:1]), s[:1])

}
