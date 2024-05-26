package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", HelloHandler)
	server.Handle("POST /blog", Blog{Title: "Title"})
	if err := http.ListenAndServe(":8080", server); err != nil {
		panic(err)
	}
}

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, `{"message": "Hello World"}`)
}

type Blog struct {
	Title string
}

func (blog Blog) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, `{"title": "`+blog.Title+`"}`)
}
