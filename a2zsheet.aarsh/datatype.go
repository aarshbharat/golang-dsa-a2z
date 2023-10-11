package a2zsheetaarsh

import "fmt"

func DataTypes(str string) (int, error) {
	if str == "Integer" || str == "Float" {
		return 4, nil
	} else if str == "Long" || str == "Double" {
		return 8, nil
	} else if str == "Character" {
		return 1, nil
	} else {
		return -1, nil
	}
	return 0, fmt.Errorf("invalid input !!!")
}