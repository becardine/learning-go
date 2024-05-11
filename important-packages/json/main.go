package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number int `json:"n"`
	Amount int `json:"a"`
}

func main() {
	account := Account{
		Number: 1001,
		Amount: 100,
	}

	// marshal converte a struct em json
	// ao usar marshal, é guardado o valor em uma variável, que pode ser usada posteriormente
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	// json é um array de bytes, então é necessário converter para string
	println(string(res))

	// encode converte a struct em json e escreve no output
	// ao usar encode, o valor é escrito diretamente no output
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	jsonPure := `{"Number":1001,"Amount":100}`
	var accountX Account
	// unmarshal converte o json em struct
	// ao usar unmarshal, é guardado o valor em uma variável, que pode ser usada posteriormente
	err = json.Unmarshal([]byte(jsonPure), &accountX)
	if err != nil {
		panic(err)
	}

	println(accountX.Number, accountX.Amount)
}
