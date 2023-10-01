package a2zsheetaarsh

import (
	"fmt"
	"unicode"
)

func FindCharCase() error {
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return err
	}

	if unicode.IsUpper(rune(input[0])) {
		fmt.Println("1")
	} else if unicode.IsLower(rune(input[0])) {
		fmt.Println("0")
	} else {
		fmt.Println("-1")
	}
	return nil
}
