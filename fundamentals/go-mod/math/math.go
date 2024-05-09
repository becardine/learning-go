package math

// para a função ficar acessível em outros pacotes, é necessário que o nome da função comece com letra maiúscula
// para que o go entenda que a função é pública
func Sum[T int | float64](a, b T) T {
	return a + b
}

var A int = 10
