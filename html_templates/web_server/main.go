package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name string
	Time int
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := tmp.Execute(writer, []Course{
			{Name: "Java", Time: 120},
			{Name: "Go", Time: 100},
			{Name: "Kotlin", Time: 85},
			{Name: "NodeJS", Time: 60},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
