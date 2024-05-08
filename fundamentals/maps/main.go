package main

import "fmt"

func main() {
	// hash table, dictionary, associative array
	salaries := map[string]float64{
		"john":   1000.0,
		"paul":   2000.0,
		"george": 3000.0,
	}

	// printando o map
	fmt.Println(salaries)

	// printando o valor de uma chave específica
	fmt.Println("salaries[\"john\"]: ", salaries["john"])

	// removendo um item do map
	delete(salaries, "paul")
	fmt.Println(salaries)

	// adicionando um item ao map
	salaries["ringo"] = 4000.0
	fmt.Println(salaries)

	// criando um map vazio
	// map[keyType]valueType
	salaries2 := make(map[string]float64) // map[string]int{} também é válido

	// adicionando itens ao map vazio
	salaries2["john"] = 1000.0
	salaries2["paul"] = 2000.0
	salaries2["george"] = 3000.0

	fmt.Println(salaries2)

	// percorrendo um map
	for key, value := range salaries2 {
		fmt.Printf("%s: %.2f\n", key, value)
	}

	// ignorando a chave
	// chamado de blank identifier, é para ignorar a chave
	for _, value := range salaries2 {
		fmt.Printf("%.2f\n", value)
	}
}
