package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Location struct {
	Type        string `json:"type"`
	Coordinates struct {
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	} `json:"coordinates"`
}

type Address struct {
	Cep          string   `json:"cep"`
	State        string   `json:"state"`
	City         string   `json:"city"`
	Neighborhood string   `json:"neighborhood"`
	Street       string   `json:"street"`
	Service      string   `json:"service"`
	Location     Location `json:"location"`
}

func main() {
	for _, postalCode := range os.Args[1:] {
		request, err := http.Get("https://brasilapi.com.br/api/cep/v2/" + postalCode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ocorreu um erro na requisição: %v\n", err)
		}
		defer request.Body.Close()
		response, err := io.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		var address Address
		if err := json.Unmarshal(response, &address); err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "Cep: %s\n\n", address)
	}
}
