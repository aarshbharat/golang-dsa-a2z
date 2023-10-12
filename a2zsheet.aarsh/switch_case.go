package a2zsheetaarsh

import (
	"fmt"
	"math"
)

func SwitchCaseExample() (float32, error) {

	var choice int8
	fmt.Println("Enter the 1 for area of circle and 2 for area of react angle")
	_, err := fmt.Scan(&choice)
	if err != nil {
		return -1, err
	}

	var r, l, b float32

	switch choice {
	case 1:
		fmt.Println("Enter the radius of circle : ")
		_, err := fmt.Scan(&r)
		if err != nil {
			return -1, err
		}
		area := float32(math.Pi * r * r)
		return area, nil
	case 2:
		fmt.Println("Enter the length :")
		_, err := fmt.Scan(&l)
		if err != nil {
			return -1, err
		}
		fmt.Println("Enter the length :")
		_, err = fmt.Scan(&b)
		if err != nil {
			return -1, err
		}
		area := l * b
		return area, nil
	}
	return -1, nil 
}
