package main

import (
	"errors"
	"fmt"
)

func main() {
	// Declarando uma função
	// func nomeDaFuncao(parametro1 tipo, parametro2 tipo) tipoDeRetorno {
	// 	// corpo da função
	// }

	// chamando a função sum
	// e printando o retorno da função
	fmt.Println(sum(1, 2))

	value, err := swap(2, 1)
	// verificando se houve erro e printando
	if err != nil {
		fmt.Println(err)
		return
	}
	// se não houve erro, printa o valor
	fmt.Println(value)
}

// Declarando uma função
// recebe dois parâmetros inteiros e retorna um inteiro
// caso a função tenha mais de um parâmetro do mesmo tipo, é possível omitir o tipo dos parâmetros
// func sum(a, b int) int {}
func sum(a, b int) int {
	return a + b
}

// uma função pode retornar mais de um valor
// func swap(a, b int) (int, int) {}
// um caso comum é retornar um valor e um error, para indicar se houve erro ou não, pois no Go não existe exceções
func swap(a, b int) (int, error) {
	if a > b {
		return 0, errors.New("a must be less than b")
	}

	// nil é um valor vazio para o tipo error
	return b + a, nil
}
