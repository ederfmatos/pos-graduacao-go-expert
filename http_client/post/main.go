package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Eder"}`))
	response, err := client.Post("https://httpbin.org/post", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)
}
