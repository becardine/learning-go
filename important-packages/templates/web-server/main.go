package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
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
		err := t.Execute(w, courses)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
