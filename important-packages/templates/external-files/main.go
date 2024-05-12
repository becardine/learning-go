package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func main() {
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

	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, courses)
	if err != nil {
		panic(err)
	}

}
