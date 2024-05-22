package main

import (
	"fmt"
	"time"
)

// go routine is a lightweight thread of execution
// go routines are multiplexed into a small number of OS threads
// go routines are managed by the go runtime
// go routines are cheap, we can have thousands of go routines in a single program

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// main é a thread principal
func main() {

	// thread 2
	go task("any")

	// thread 3
	go task("another")

	// thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task anonymous is running\n", i)
			time.Sleep(1 * time.Second)
		}
	}()

	fmt.Println("Waiting for go routines to finish")
	time.Sleep(15 * time.Second)
}
