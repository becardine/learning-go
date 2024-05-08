package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

// variadic function
// aceita um número variável de argumentos, que precisam ser do mesmo tipo
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
