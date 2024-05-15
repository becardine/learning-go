package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request initialized")
	defer log.Println("request finalized")
	select {
	case <-time.After(5 * time.Second):
		// imprime no console que a requisição foi processada com sucesso
		log.Println("processing request successfully")
		// imprime no browser que a requisição foi processada com sucesso
		w.Write([]byte("processing request successfully"))
	case <-ctx.Done():
		err := ctx.Err()
		log.Println("request cancelled:", err)
	}
}
