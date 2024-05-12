package main

import "net/http"

func main() {
	// cria um novo servidor mux
	// serve multiplos servidores,
	// ao criar um novo mux, você consegue ter mais controle sobre as rotas e handlers
	mux := http.NewServeMux()

	// adiciona um handler para a rota /
	// passando uma função anônima
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, World!"))
	// })

	// passando uma função nomeada
	mux.HandleFunc("/", HomeHandler)

	// passando uma struct que implementa a interface ServeHTTP
	// ao chamar a rota /blog, o método ServeHTTP da struct blog será chamado
	mux.Handle("/blog", blog{title: "Blog Page"})

	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page"))
}

// cria uma struct blog
type blog struct {
	title string
}

// implementa a interface ServeHTTP para a struct blog
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
