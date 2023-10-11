package main

import (
	"fmt"

	a2zsheetaarsh "a2zsheet.aarsh/a2zsheet.aarsh"
)

func main() {
	// Call for FindCharCase method
	a2zsheetaarsh.FindCharCase()

	// Call for DataTypes method
	output, err := a2zsheetaarsh.DataTypes("Integer")
	fmt.Println("Output of function a2zsheetaarsh.DataTypes : ", output, "error : ", err)
}
