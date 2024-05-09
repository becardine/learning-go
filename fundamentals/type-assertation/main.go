package main

func main() {
	// type assertation
	// é usado para verificar se um valor de interface é de um tipo específico
	// é usado para obter o valor subjacente de uma interface

	var a interface{} = "be cardine"
	println(a)          // (0x4c7040,0x4eca90)
	println(a.(string)) // be cardine

	// caso o valor não seja do tipo esperado, o programa irá gerar um erro
	res, ok := a.(int)
	println(res, ok) // 0 false
}
