package main

func main() {
	// deadlock
	// ch := make(chan string)

	// com o buffer de 2, o código não entra em deadlock
	// no entanto, não é recomendado usar pois vai estar usando memória desnecessariamente
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
