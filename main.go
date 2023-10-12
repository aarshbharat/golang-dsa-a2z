package main

import (
	"fmt"

	a2zsheetaarsh "a2zsheet.aarsh/a2zsheet.aarsh"
)

func main() {
	// Call for FindCharCase method
	a2zsheetaarsh.FindCharCase()

	// Call for DataTypes method
	output1, err := a2zsheetaarsh.DataTypes("Integer")
	fmt.Println("Output of function a2zsheetaarsh.DataTypes :", output1, "error : ", err)

	// Call for IntegerComparision method
	output2, err := a2zsheetaarsh.IntegerComparision(int(5), int(4))
	fmt.Println("Output of function a2zsheetaarsh.IntegerComparision :e", output2, "error : ", err)

	// Switch case example 
	output3,_ := a2zsheetaarsh.SwitchCaseExample()
	fmt.Println("Area : ", output3)

}
