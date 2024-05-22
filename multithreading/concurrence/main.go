package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		// soma atimica para evitar concorrencia
		atomic.AddUint64(&number, 1)
		// m.Unlock()
		w.Write([]byte(fmt.Sprintf("You are visitor number %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}
