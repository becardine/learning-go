package main

func main() {
	forever := make(chan bool) // canal está vazio
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true // envia um valor para o canal
	}()
	// deadlock: o canal está vazio e não há ninguém para enviar um valor
	<-forever // espera por um valor no canal

}
