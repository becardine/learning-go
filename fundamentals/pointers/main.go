package main

import "fmt"

func sum(a, b *int) int {
	*a = 15
	return *a + *b
}

func main() {
	// ponteiros são tipos de dados que armazenam endereços de memória
	// para criar um ponteiro, basta usar o operador "&" seguido do valor que se deseja armazenar
	// variavel aponta para um ponteiro que tem o endereço da memória onde o valor está armazenado
	a := 10
	fmt.Println(&a) // 0xc0000140a8

	// o ponteiro é um armazenamento na memória
	var pointer *int = &a
	fmt.Println(pointer) // 0xc0000140a8

	// ao atribuir um valor a um ponteiro, o valor é armazenado no endereço de memória que o ponteiro aponta
	*pointer = 20
	fmt.Println(a) // 20

	b := &a
	fmt.Println(b)  // 0xc0000140a8
	fmt.Println(*b) // 20

	var1 := 10
	var2 := 20
	// ao passar um valor para uma função, o valor é copiado para a função
	// caso seja alterado esse valor, a alteração não será refletida na variável original
	// fmt.Println(sum(var1, var2)) // 30

	// para passar o endereço de memória de uma variável para uma função, basta passar o ponteiro
	// assim, a função pode alterar o valor da variável original
	fmt.Println(sum(&var1, &var2)) // 30

	// quando usar ponteiros?
	// - quando for necessário alterar o valor de uma variável dentro de uma função
	// - quando quiser que os valores sejam mutáveis

	// quando não usar ponteiros?
	// - quando não for necessário alterar o valor de uma variável dentro de uma função
}
