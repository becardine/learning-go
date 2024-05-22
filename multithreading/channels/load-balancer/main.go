package main

import "time"

func worker(workerId int, jobs <-chan int) {
	for j := range jobs {
		println("Worker", workerId, "started job", j)
		time.Sleep(time.Second)
	}
}

func main() {
	// declara o canal
	jobs := make(chan int)
	// cria 3 workers
	workersQty := 10
	// inicia os workers
	for w := 1; w <= workersQty; w++ {
		go worker(w, jobs)
	}

	for j := 1; j <= 100; j++ {
		jobs <- j
	}

	close(jobs)
	time.Sleep(time.Second)
}
