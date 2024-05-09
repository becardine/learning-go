package main

import "fmt"

type Address struct {
	Street  string
	Number  int
	Country string
}

type Client struct {
	Name    string
	Age     int
	Active  bool
	Address Address
}

// é possível criar métodos para structs
// para isso, basta definir uma função com um receiver do tipo da struct
// o receiver é um parâmetro que permite acessar os campos da struct
// o receiver é definido entre parênteses antes do nome da função
func (client *Client) Deactivate() {
	client.Active = false
	fmt.Println("Client deactivated")
}

// interfaces são tipos de dados que definem um conjunto de métodos, não pode conter propriedades (strigs, int, etc)
// structs podem implementar interfaces
type Person interface {
	Deactivate()
}

func DeactivatePerson(person Person) {
	person.Deactivate()
}

func main() {
	// go não é uma linguagem orientada a objetos, mas é possível criar structs que se comportam de forma semelhante a classes
	// structs são tipos de dados que podem conter campos de diferentes tipos
	// para criar uma struct, basta usar a palavra reservada "type" seguida do nome da struct e a palavra reservada "struct"
	// os campos da struct são definidos entre chaves
	// é possível criar structs aninhadas, ou seja, structs que contém outras structs
	// go way of developing: structs

	client := Client{
		Name:   "Maria",
		Age:    25,
		Active: true,
	}

	// para acessar um campo de uma struct, basta usar o operador ponto
	// client.Active = false

	// é possível acessar os campos de structs aninhadas da mesma forma
	client.Address = Address{
		Street:  "Rua 1",
		Number:  123,
		Country: "Brasil",
	}

	fmt.Println(client) // {Maria 25 true {Rua 1 123 Brasil}}

	// client.Deactivate() // Client deactivated
	// fmt.Println(client) // {Maria 25 false {Rua 1 123 Brasil}}

	DeactivatePerson(&client) // Client deactivated
	fmt.Println(client)       // {Maria 25 false {Rua 1 123 Brasil}}
}
