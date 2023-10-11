package a2zsheetaarsh

import "fmt"

func IntegerComparision(a, b int) (string, error) {
	if a > b {
		return "greater", nil
	} else if a < b {
		return "smaller", nil
	} else {
		return "equal", nil
	}
	return "", fmt.Errorf("invalid input !!!!")
}