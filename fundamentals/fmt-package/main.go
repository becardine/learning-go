package main

// ao importar um pacote, é necessário utilizá-lo
// caso contrário, o código não compila
import "fmt"

// quando tem mais de um pacote, pode ser importado desta forma
// import (
// 	"fmt"
// 	"math"
// )

var x float64 = 1.8

func main() {
	// %T é um verb que retorna o tipo da variável
	// para retornar o valor da variável, é necessário utilizar o verb %v
	fmt.Printf("O valor de X é %T", x)
}