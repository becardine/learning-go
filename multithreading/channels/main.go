package main

import "fmt"

// thread 1
func main() {
	channel := make(chan string)
	// thread 2
	go func() {
		channel <- "Hello World!"
	}()
	// o valor é recebido do canal é armazenado na variável message
	message := <-channel
	fmt.Println(message)
}
