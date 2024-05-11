package main

import (
	"encoding/json"
	"io"
	"net/http"
)

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
	// subindo um servidor http
	http.HandleFunc("/search-cep", SearchCepHandler)
	http.ListenAndServe(":8080", nil)
}

func SearchCepHandler(w http.ResponseWriter, r *http.Request) {
	// recuperando o cep da query
	cep := r.URL.Query().Get("cep")

	// verificando se o cep foi passado
	if cep == "" {
		// retornando um erro caso o cep não tenha sido passado
		http.Error(w, "cep is required", http.StatusBadRequest)
		return
	}

	// buscando o cep
	viaCep, err := SearchCep(cep)
	if err != nil {
		// retornando um erro caso ocorra um erro na busca do cep
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// setando o content type da resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// retornando o cep encontrado
	json.NewEncoder(w).Encode(viaCep)
}

func SearchCep(cep string) (*ViaCep, error) {
	res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	viaCep := ViaCep{}
	err = json.Unmarshal(body, &viaCep)
	if err != nil {
		return nil, err
	}

	return &viaCep, nil
}
