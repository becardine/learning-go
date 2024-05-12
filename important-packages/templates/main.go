package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

func main() {
	golang := Course{
		Name:  "Golang",
		Hours: 20,
	}

	tmp := template.New("CourseTemplate")
	tmp, _ = tmp.Parse("Course Name: {{.Name}} Hours: {{.Hours}}")
	err := tmp.Execute(os.Stdout, golang)
	if err != nil {
		panic(err)
	}

}
