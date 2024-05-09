package main

// go mod init {module-name}
// geralmente o nome do módulo é o nome do repositório

import (
	"github.com/becardine/learning-go/math"
)

func main() {

	var s = math.Sum(3, 3)
	println(s) // 6

	println(math.A) // 10
}
