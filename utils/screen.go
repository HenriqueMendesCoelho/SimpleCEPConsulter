package utils

import (
	"fmt"
)

func AskToContinue() bool {
	var con string
	fmt.Println("Want another CEP? (y/n)")
	fmt.Scanf("%s", &con)
	fmt.Scanln()

	fmt.Println()
	fmt.Println()

	return con == "y"
}

func ShowError(err error) {
	fmt.Println()
	fmt.Println("***************")
	fmt.Println("Error: ", err)
	fmt.Println("***************")
	fmt.Println()

	AskToContinue()
}
