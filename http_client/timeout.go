package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Second * 1,
	}
	response, err := client.Get("https://httpbin.org/delay/2")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
