package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name string
	Time int
}

func main() {
	course := Course{
		Name: "Go Expert",
		Time: 40,
	}
	tmp := template.Must(template.New("course").Parse("Curso: {{.Name}} - Carga Horária: {{.Time}}"))
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
