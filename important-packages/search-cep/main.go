package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// https://mholt.github.io/json-to-go/ to convert JSON to Go struct
type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	// Search CEP in Go
	for _, url := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + url + "/json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "search-cep: %v\n", err)
			continue
		}

		// defer is used to delay the execution of a function until the end of the current block
		defer req.Body.Close()

		// read the response body
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "search-cep: %v\n", err)
			continue
		}

		var viaCep ViaCep
		// unmarshal converts the json into struct
		err = json.Unmarshal(res, &viaCep)
		if err != nil {
			fmt.Fprintf(os.Stderr, "search-cep: %v\n", err)
			continue
		}

		fmt.Printf("CEP: %s\n", viaCep)

		// create a file to store the CEP information
		file, err := os.Create("cep.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "search-cep: %v\n", err)
			continue
		}

		// close the file after writing to release resources
		defer file.Close()

		// write the CEP information to the file
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Logradouro: %s, Complemento: %s, Bairro: %s, Localidade: %s, UF: %s, IBGE: %s, GIA: %s, DDD: %s, SIAFI: %s\n", viaCep.Cep, viaCep.Logradouro, viaCep.Complemento, viaCep.Bairro, viaCep.Localidade, viaCep.Uf, viaCep.Ibge, viaCep.Gia, viaCep.Ddd, viaCep.Siafi))
	}
}
