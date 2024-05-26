package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", FindAddressHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	fmt.Printf("Server running")
}

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

func FindAddressHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	postalCode := request.URL.Query().Get("postalCode")
	if postalCode == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, `{"error": "Informe o cep"}`)
		return
	}
	writer.WriteHeader(http.StatusOK)
	address, err := FindAddress(postalCode)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, `{"error": "`+err.Error()+`"}`)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(address)
}

func FindAddress(postalCode string) (*Address, error) {
	request, err := http.Get("https://brasilapi.com.br/api/cep/v2/" + postalCode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ocorreu um erro na requisição: %v\n", err)
		return nil, err
	}
	defer request.Body.Close()
	response, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	var address Address
	if err := json.Unmarshal(response, &address); err != nil {
		return nil, err
	}
	return &address, nil
}
