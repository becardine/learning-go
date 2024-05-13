package main

// quando o template for html, o pacote recomendado é o html/template
// existe algumas restrições de segurança
import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	// ... significa que é uma função variádica, ou seja, pode receber um número variável de argumentos
	// t := template.Must(template.New("content.html").ParseFiles(templates...))

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	courses := Courses{
		{
			Name:  "Golang",
			Hours: 20,
		},
		{
			Name:  "Python",
			Hours: 30,
		},
		{
			Name:  "Java",
			Hours: 40,
		},
	}

	err := t.Execute(os.Stdout, courses)
	if err != nil {
		panic(err)
	}

}
