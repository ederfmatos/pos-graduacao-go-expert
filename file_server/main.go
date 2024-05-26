package main

import (
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./file_server/public"))
	server := http.NewServeMux()
	server.Handle("/", fileServer)
	_ = http.ListenAndServe(":8080", server)
}
