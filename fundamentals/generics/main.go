package main

// constraints são usadas para restringir os tipos de dados que podem ser usados em uma função genérica
type Number interface {
	// ao usar o ~, o go entende que pode usar o tipo MyNumber
	~int | float64
}

type MyNumber int

func sum[T Number](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

func compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	// generics é um recurso que permite escrever funções, tipos e estruturas de dados que trabalham com tipos arbitrários
	// é uma maneira de escrever código que é independente do tipo de dados

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	println(sum(m)) // 6

	n := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	println(sum(n)) // +6.600000e+000

	o := map[string]MyNumber{"a": 1, "b": 2, "c": 3}
	println(sum(o)) // 6

	println(compare(1, 1))     // true
	println(compare(1, 2))     // false
	println(compare(10, 10.0)) // true
}
