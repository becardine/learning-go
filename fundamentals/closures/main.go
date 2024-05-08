package main

import "fmt"

func main() {
	// closure é uma função que captura variáveis do escopo em que foi declarada
	// e que pode ser executada em qualquer lugar
	// é uma função anônima, que é atribuída a uma variável ou passada como argumento para outra função

	// um exemplo de closure é utilizar para rodar um web server

	result := func() int {
		// aqui pode ter outras funções, loops, ifs, etc
		return 42 * 7
	}()

	fmt.Println(result)
}
