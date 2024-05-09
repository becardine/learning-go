package main

import "fmt"

type Client struct {
	Name string
}

func (c *Client) changeName(name string) {
	c.Name = name
}

func main() {

	client := Client{Name: "John Doe"}
	fmt.Println(client) // {John Doe}

	client.changeName("Jane Doe")
	fmt.Println(client) // {Jane Doe}
}
