package main

import (
	"log"
	"net/http"
)

func main() {
	// cria um servidor de arquivos
	// serve arquivos estáticos
	fileServer := http.FileServer(http.Dir("./public"))

	mux := http.NewServeMux()

	// adiciona o servidor de arquivos ao servidor mux
	mux.Handle("/", fileServer)

	// adiciona um handler para a rota /home
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
