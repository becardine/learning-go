package main

import (
	"fmt"

	"github.com/becardine/learning-go/packaging/math"
)

// folder cmd is used to store the main package

func main() {
	// m := math.Math{A: 1, B: 2}
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
