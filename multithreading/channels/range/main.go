package main

import "fmt"

func main() {
	channel := make(chan int)
	go publish(channel)
	reader(channel) // esse metodo não pode ser executado em background, pois ele é responsável por imprimir os valores do canal e esvaziar o canal
}

func publish(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel) // ao fechar o canal, o método reader não entrará em deadlock
}

func reader(channel chan int) {
	for i := range channel {
		fmt.Printf("Received: %d\n", i)
	}

}
