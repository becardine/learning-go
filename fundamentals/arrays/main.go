package main

import "fmt"

func main() {
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5
	fmt.Println(len(array) - 1)

}