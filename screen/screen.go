package screen

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"com.github/henriquemendescoelho/simplecepconsultor/service"
)

func AskToContinue() bool {
	var con string
	fmt.Println("Buscar outro CEP? (s/n)")
	fmt.Scanf("%s", &con)
	fmt.Scanln()

	fmt.Println()
	fmt.Println()

	return strings.ToLower(con) == "s" || strings.ToLower(con) == "sim"
}

func ShowError(err error) {
	fmt.Println()
	fmt.Println("***************")
	fmt.Println("Erro: ", err)
	fmt.Println("***************")
	fmt.Println()
}

func GenerateMessage(res service.CepResponse) string {
	cep := fmt.Sprintf("Cep: %s", res.Cep)
	logradouro := fmt.Sprintf("Logradouro: %s", res.Logradouro)
	bairro := fmt.Sprintf("Bairro: %s", res.Bairro)
	localidade := fmt.Sprintf("Localidade: %s", res.Localidade)
	uf := fmt.Sprintf("UF: %s", res.Uf)
	ddd := fmt.Sprintf("DDD: %s", res.Ddd)

	return strings.Join([]string{cep, logradouro, bairro, localidade, uf, ddd}, "\n")
}

func AskCEP() string {
	var cep string

	fmt.Println("Insira o CEP: (apenas números)")
	fmt.Scanf("%s", &cep)
	fmt.Scanln()

	fmt.Println()
	fmt.Println("=================")
	fmt.Println()

	return cep
}

func ClearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
