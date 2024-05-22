package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)              // 10 goroutines
	go publish(channel)     // publica o valor no canal
	go reader(channel, &wg) // lê o valor do canal, nesse caso ela pode ser executada em background utilizando wait group
	wg.Wait()               // espera por todas as goroutines
}

func publish(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

func reader(channel chan int, wg *sync.WaitGroup) {
	for i := range channel {
		fmt.Printf("Received: %d\n", i)
		wg.Done()
	}
}
