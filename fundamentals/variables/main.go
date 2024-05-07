package main

const a = "Hello, World!"

// variáveis globais são inicializadas com valores padrão
var (
	b bool    // infere o valor false
	c int     // infere o valor 0
	d string  // infere o valor ""
	e float64 // infere o valor +0.000000e+000
)

func main() {
	// fortemente tipada, não é possível atribuir um valor de outro tipo
	// b = 10

	// variáveis locais que não são utilizadas geram erro de compilação
	// var a string = "X"

	// variáveis locais
	// := é uma forma de declarar e inicializar uma variável local (short variable declaration)
	// := só pode ser utilizada quando uma variável é declarada pela primeira vez
	c := "X"
	println(c)
}