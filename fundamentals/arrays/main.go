package main

import "fmt"

func main() {
	// Declarando um array de 5 posições, não é possível adicionar mais posições
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	fmt.Println(len(array) - 1)
	fmt.Println(array[0])

	// Percorrendo o array
	for i, v := range array {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
