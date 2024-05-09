package main

func main() {
	// for é a única estrutura de repetição em go
	for i := 0; i < 10; i++ {
		println(i)
	}

	numbers := []int{1, 2, 3, 4, 5}
	// range retorna o índice e o valor
	// se você não precisa de um dos valores, utilize _ para ignorá-lo
	for i, n := range numbers {
		println(i, n)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// loop infinito
	// geralmente para consumir mensagens de uma fila ou para rodar um servidor
	// for {
	// 	println("loop")
	// }
}
