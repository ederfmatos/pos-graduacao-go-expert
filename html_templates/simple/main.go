package main

import (
	"log"
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
	tmp := template.New("course")
	tmp, err := tmp.Parse("Curso: {{.Name}} - Carga Horária: {{.Time}}")
	if err != nil {
		panic(err)
	}
	err = tmp.Execute(os.Stdout, course)
	log.Fatalf("Error at execute template: %v", err)
}
