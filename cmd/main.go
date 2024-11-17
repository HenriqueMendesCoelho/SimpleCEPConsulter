package main

import (
	"fmt"
	"os"

	"com.github/henriquemendescoelho/simplecepconsultor/screen"
	"com.github/henriquemendescoelho/simplecepconsultor/service"
)

func main() {
	var cep string = screen.AskCEP()

	screen.ClearScreen()

	result, err := service.GetCep(cep)
	if err == nil {
		fmt.Println(screen.GenerateMessage(result))
		fmt.Println()
	} else {
		screen.ShowError(err)
	}

	if screen.AskToContinue() {
		screen.ClearScreen()
		main()
	}

	os.Exit(0)
}
