package main

// go mod init {module-name}
// geralmente o nome do módulo é o nome do repositório
// go mod tidy - remove dependências que não estão sendo utilizadas e adiciona dependências que estão sendo utilizadas

import (
	"github.com/becardine/learning-go/math"
)

func main() {

	var s = math.Sum(3, 3)
	println(s) // 6

	println(math.A) // 10
}
