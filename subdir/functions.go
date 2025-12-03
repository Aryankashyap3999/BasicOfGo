package subdir

import "fmt"

func Function (num int) (string, int) {
	if num % 2 == 0 {
		fmt.Println("Even")
		return "Even", 0
	} else {
		fmt.Println("odd")
		return "Odd", 1
	}
}