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
	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := tmp.Execute(os.Stdout, []Course{
		{Name: "Java", Time: 120},
		{Name: "Go", Time: 100},
		{Name: "Kotlin", Time: 85},
		{Name: "NodeJS", Time: 60},
	})
	if err != nil {
		panic(err)
	}
}
