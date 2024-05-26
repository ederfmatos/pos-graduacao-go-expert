package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number  int     `json:"number,omitempty"`
	Balance float64 `json:"balance,omitempty"`
}

func (account Account) String() string {
	return `Account["number":` + fmt.Sprint(account.Number) + `,"balance":` + fmt.Sprint(account.Balance) + `]`
}

func main() {
	account := Account{
		Number:  1,
		Balance: 100,
	}
	contentBytes, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(contentBytes))

	// Encoder
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(account)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("account.json")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	encoder = json.NewEncoder(file)
	err = encoder.Encode(account)
	if err != nil {
		panic(err)
	}
	err = os.Remove("account.json")
	if err != nil {
		panic(err)
	}

	accountJson := []byte(`{"number":1,"balance":100}`)
	var myAccount Account
	err = json.Unmarshal(accountJson, &myAccount)
	if err != nil {
		panic(err)
	}
	fmt.Println(myAccount)
}
