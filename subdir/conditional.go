package subdir

import "fmt"

func Conditional () {
	fmt.Println("Using Conditionals")

	var number = 2

	if number % 2 == 0 {
		fmt.Println("Even number")
	} else if number % 2 != 0 {
		fmt.Println("Odd number")
	} else {
		fmt.Println("Invalid number")
	}
}