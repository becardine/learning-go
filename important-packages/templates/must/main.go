package main

import (
	"html/template"
	"os"
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

	tmp := template.Must(template.New("CourseTemplate").Parse("Course Name: {{.Name}} Hours: {{.Hours}}"))
	err := tmp.Execute(os.Stdout, golang)
	if err != nil {
		panic(err)
	}
}
