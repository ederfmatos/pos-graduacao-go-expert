package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	request, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("AnotherHeader", "Random Value")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)
}
