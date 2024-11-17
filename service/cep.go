package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type CepResponse struct {
	Cep        string
	Logradouro string
	Bairro     string
	Localidade string
	Uf         string
	Ddd        string
	Erro       bool
}

const notFounderrorMessage = "CEP não encontrado, tente novamente"
const formatErrorMessage = "CEP deve ter 8 digitos, tente novamente"
const errorMessage = "erro ao buscar CEP"

func GetCep(cep string) (CepResponse, error) {
	var cepResp CepResponse
	const viaCEPUrl = "https://viacep.com.br/ws/%s/json/"
	url := fmt.Sprintf(viaCEPUrl, cep)

	resp, err := http.Get(url)
	if err != nil {
		return cepResp, errors.New(notFounderrorMessage)
	}
	if resp.StatusCode == http.StatusBadRequest {
		return cepResp, errors.New(formatErrorMessage)
	}
	if resp.StatusCode != http.StatusOK {
		return cepResp, errors.New(errorMessage)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return cepResp, err
	}

	err = json.Unmarshal(body, &cepResp)
	if err != nil {
		return cepResp, errors.New(notFounderrorMessage)
	}

	if cepResp.Erro {
		return cepResp, errors.New(notFounderrorMessage)
	}

	return cepResp, nil
}
