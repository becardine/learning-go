package main

import "fmt"

func main() {
	// empty interface
	// é uma interface que não possui métodos
	// pode ser usada para armazenar valores de qualquer tipo
	// é usada quando não se sabe o tipo de dado que será armazenado
	// deve ser evitada, pois pode tornar o código menos legível e mais difícil de manter
	// e o go é uma linguagem fortemente tipada
	// generics pode ser uma solução melhor

	var a interface{} = 10
	var y interface{} = "Hello, World!"
	fmt.Println(a) // 10

	showType(a) // 10
	showType(y) // string
}

func showType(x interface{}) {
	fmt.Printf("Type: %T\n", x)
}
