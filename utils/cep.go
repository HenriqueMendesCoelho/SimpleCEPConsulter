package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func GetCep(cep string) (CepResponse, error) {
	var cepResp CepResponse

	baseUrl := "https://viacep.com.br/ws/%s/json/"
	url := fmt.Sprintf(baseUrl, cep)

	resp, err := http.Get(url)
	if err != nil {
		return cepResp, errors.New("cep not found")
	}

	if resp.StatusCode == http.StatusBadRequest {
		return cepResp, errors.New("cep must contain only numbers and length must be 8, try again")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return cepResp, err
	}

	err = json.Unmarshal(body, &cepResp)
	if err != nil {
		return cepResp, err
	}

	if cepResp.Erro {
		return cepResp, errors.New("cep not found")
	}

	return cepResp, nil
}

func GenerateMessage(res CepResponse) string {
	cep := fmt.Sprintf("Cep: %s", res.Cep)
	logradouro := fmt.Sprintf("Logradouro: %s", res.Logradouro)
	bairro := fmt.Sprintf("Bairro: %s", res.Bairro)
	localidade := fmt.Sprintf("Localidade: %s", res.Localidade)
	uf := fmt.Sprintf("UF: %s", res.Uf)
	ddd := fmt.Sprintf("DDD: %s", res.Ddd)

	return strings.Join([]string{cep, logradouro, bairro, localidade, uf, ddd}, "\n")
}
