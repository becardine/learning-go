package main

// chan<- significa que o canal só pode receber valores - receive only
func receive(name string, hello chan<- string) {
	hello <- name
}

// <-chan significa que o canal só pode enviar valores - send only
func read(data <-chan string) {
	name := <-data
	println("Hello, " + name)
}

func main() {
	hello := make(chan string, 1)
	go receive("World", hello)
	read(hello)

}
