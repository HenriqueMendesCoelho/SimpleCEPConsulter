package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"com.github/henriquemendescoelho/simplecepconsultor/utils"
)

func main() {
	var cep string

	fmt.Println("Enter CEP to search: (only numbers)")
	fmt.Scanf("%s", &cep)
	fmt.Scanln()

	fmt.Println()
	fmt.Println("=================")
	fmt.Println()

	ClearScreen()

	result, err := utils.GetCep(cep)
	if err != nil {
		utils.ShowError(err)
		utils.AskToContinue()
	}

	fmt.Println(utils.GenerateMessage(result))

	fmt.Println()

	if utils.AskToContinue() {
		ClearScreen()
		main()
	}

	os.Exit(0)
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
